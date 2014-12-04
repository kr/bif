package interp

import (
	"bufio"
	"fmt"
	"go/parser"
	"os"
)

func Run(env *Env, packages map[string]Package) {
	scanner := bufio.NewScanner(os.Stdin)
	for prompt("; ", scanner) {
		x, err := parser.ParseExpr(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		a, err := Eval(x, env)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, v := range a {
			fmt.Printf("%+v\n", v.Interface())
		}
	}
	fmt.Println()
}
