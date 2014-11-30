package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

var wd, _ = os.Getwd()

func main() {
	tmpdir, err := ioutil.TempDir("", "bif")
	if err != nil {
		log.Fatalln(err)
	}
	defer os.RemoveAll(tmpdir)
	gopath := tmpdir
	if s := os.Getenv("GOPATH"); s != "" {
		gopath += string(os.PathListSeparator) + s
	}

	p, err := build.ImportDir(wd, 0)
	if err != nil {
		log.Fatalln(err)
	}

	rel, _ := filepath.Rel(p.Root, p.Dir)
	pdir := filepath.Join(tmpdir, rel)
	os.MkdirAll(pdir, 0777)
	for _, name := range p.GoFiles {
		err = copyfile(name, filepath.Join(pdir, name))
		if err != nil {
			log.Fatalln(err)
		}
	}
	err = gen(filepath.Join(pdir, "bif.go"), p)
	if err != nil {
		log.Fatalln(err)
	}
	bifmain := filepath.Join(tmpdir, "bifmain.go")
	err = genbifmain(bifmain, p)
	if err != nil {
		log.Fatalln(err)
	}
	//arg := append([]string{"run", "bif.go"}, p.GoFiles...)
	//cmd := exec.Command("go", arg...)
	cmd := exec.Command("go", "run", bifmain)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(envNoGopath(), "GOPATH="+gopath)
	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func envNoGopath() (a []string) {
	for _, s := range os.Environ() {
		if !strings.HasPrefix(s, "GOPATH=") {
			a = append(a, s)
		}
	}
	return a
}

func copyfile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()
	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, r)
	if err != nil {
		w.Close()
		return err
	}
	return w.Close()
}

var (
	pmainTmpl = template.Must(template.New("main").Parse(`
package main

import (
	"os"
	"reflect"

	"github.com/kr/bif/interp"
)

func BifMain() {
	var env *interp.Env
{{range .Globals}}
	env = env.Augment({{.Name|printf "%q"}}, reflect.ValueOf({{.Name}}))
{{end}}

	// TODO(kr): packages
	interp.Run(env, nil)
	os.Exit(0)
}
`))
	bifmainTmpl = template.Must(template.New("main").Parse(`
package main

import p {{.ImportPath|printf "%q"}}

func main() {
	p.BifMain()
}
`))
)

func gen(name string, p *build.Package) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	tab, _, err := syms(p, true)
	if err != nil {
		return err
	}

	var x = struct{ Globals []sym }{tab}

	if err := pmainTmpl.Execute(f, x); err != nil {
		return err
	}
	return nil
}

func genbifmain(name string, p *build.Package) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	var x = struct{ ImportPath string }{p.ImportPath}
	if err := bifmainTmpl.Execute(f, x); err != nil {
		return err
	}
	return nil
}

func syms(p *build.Package, includePrivate bool) ([]sym, map[string][]string, error) {
	fset := token.NewFileSet() // positions are relative to fset

	pkgs, err := parser.ParseDir(fset, ".", func(fi os.FileInfo) bool {
		for _, s := range p.GoFiles {
			if s == fi.Name() {
				return true
			}
		}
		return false
	}, 0)
	if err != nil {
		return nil, nil, err
	}

	pkg, ok := pkgs[p.Name]
	if !ok {
		return nil, nil, fmt.Errorf("no package %s here", p.Name)
	}

	if !includePrivate {
		ok = ast.PackageExports(pkg)
		if !ok {
			return nil, nil, fmt.Errorf("no exported symbols in %s", p.Name)
		}
	}

	vis := &visitor{}
	ast.Walk(vis, pkg)

	return vis.tab, nil, nil
}

type sym struct {
	Name    string
	CanAddr bool
}

type visitor struct {
	tab []sym
}

func (v *visitor) Visit(x ast.Node) (w ast.Visitor) {
	switch x := x.(type) {
	case *ast.FuncDecl:
		if x.Recv == nil && x.Name.Name != "init" {
			v.tab = append(v.tab, sym{x.Name.Name, false})
		}
		return nil
	case *ast.ImportSpec:
		return nil
	case *ast.ValueSpec:
		// TODO(kr): keep track of writable symbols (vars)
		// and make their reflect.Value values settable
		// in pmainTmpl above.
		for _, id := range x.Names {
			v.tab = append(v.tab, sym{id.Name, false})
		}
		return nil
	case *ast.TypeSpec:
		return nil
	}
	return v
}
