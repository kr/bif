// +build ignore

/*

Generate reflected operations.
Run with:

	go run genop.go | gofmt >op.go

*/
package main

import (
	"go/token"
	"os"
	"strings"
	"text/template"
)

var (
	floatkinds = []string{
		"Float32",
		"Float64",
		"Complex64",
		"Complex128",
	}
	intkinds = []string{
		"Uint8",
		"Uint16",
		"Uint32",
		"Uint64",
		"Int8",
		"Int16",
		"Int32",
		"Int64",
		"Uint",
		"Int",
		"Uintptr",
	}

	floatops = []token.Token{
		token.ADD,
		token.SUB,
		token.MUL,
		token.QUO,
	}
	intops = []token.Token{
		token.REM,
		token.AND,
		token.OR,
		token.XOR,
		token.AND_NOT,
	}

	opnames = map[token.Token]string{
		token.ADD:     "ADD",
		token.SUB:     "SUB",
		token.MUL:     "MUL",
		token.QUO:     "QUO",
		token.REM:     "REM",
		token.AND:     "AND",
		token.OR:      "OR",
		token.XOR:     "XOR",
		token.AND_NOT: "AND_NOT",
	}
)

type ent struct {
	Op   string // language operator, eg + - *
	Tok  string // token.Token value, eg ADD SUB MUL
	Kind string // reflect.Kind value, eg Int Int32 Uint
	Type string // basic data type, eg int int32 uint
	Func string // name of generated func, eg addint mulfloat32
}

func newent(op token.Token, kind string) ent {
	return ent{
		Op:   op.String(),
		Tok:  opnames[op],
		Kind: kind,
		Type: strings.ToLower(kind),
		Func: strings.ToLower(opnames[op]) + kind,
	}
}

func main() {
	var x struct {
		Ents  []ent
		Ops   []string
		Kinds []string
	}

	x.Kinds = append(floatkinds, intkinds...)
	allops := append(floatops, intops...)
	for _, op := range allops {
		x.Ops = append(x.Ops, strings.ToUpper(opnames[op]))
	}

	for _, kind := range floatkinds {
		for _, op := range floatops {
			x.Ents = append(x.Ents, newent(op, kind))
		}
	}
	for _, kind := range intkinds {
		for _, op := range allops {
			x.Ents = append(x.Ents, newent(op, kind))
		}
	}

	binopTmpl.Execute(os.Stdout, x)
}

var binopTmpl = template.Must(template.New("op.go").Funcs(template.FuncMap{
	"tolower": strings.ToLower,
}).Parse(`
package interp

import (
	"go/token"
	"reflect"
)

type binop func(x, y interface{}) interface{}

var binops = map[token.Token]map[reflect.Kind]binop{
	{{range .Ops}}token.{{.}}: {},
	{{end}}
}

var kindtypes = map[reflect.Kind]reflect.Type{
	reflect.String: reflect.TypeOf(""),
	{{range .Kinds}}reflect.{{.}}: reflect.TypeOf({{. | tolower}}(0)),
	{{end}}
}

func init() {
	binops[token.ADD][reflect.String] = addString
	{{range .Ents}}binops[token.{{.Tok}}][reflect.{{.Kind}}] = {{.Func}}
	{{end}}
}

func addString(x, y interface{}) interface{} {
	return x.(string) + y.(string)
}

{{range .Ents}}
func {{.Func}}(x, y interface{}) interface{} {
	return x.({{.Type}}) {{.Op}} y.({{.Type}})
}
{{end}}
`))
