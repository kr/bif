package interp

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)

type Env struct {
	s string
	v reflect.Value
	e *Env
}

// Env0 returns the empty environment.
func Env0() *Env {
	return nil
}

func (e *Env) get(name string) reflect.Value {
	if e == nil {
		panic(errors.New("not found: " + name))
	}
	if e.s == name {
		return e.v
	}
	return e.e.get(name)
}

func (e *Env) Augment(name string, v reflect.Value) *Env {
	return &Env{name, v, e}
}

func catch(err *error) {
	if v := recover(); v != nil {
		if e, ok := v.(error); ok {
			*err = e
		} else {
			*err = fmt.Errorf("%v", v)
		}
	}
}

func Eval(x ast.Expr, e *Env) (vals []reflect.Value, err error) {
	defer catch(&err)
	vals = eval(x, e)
	return vals, nil
}

func Eval1(x ast.Expr, e *Env) (v reflect.Value, err error) {
	defer catch(&err)
	v = eval1(x, e)
	return v, nil
}

// eval evaluates x in e
// and returns the resulting values.
func eval(x ast.Expr, e *Env) []reflect.Value {
	var v reflect.Value
	switch x := x.(type) {
	case *ast.BasicLit:
		v = evalBasicLit(x)
	case *ast.Ident:
		v = e.get(x.Name)
	case *ast.CallExpr:
		return evalCall(x, e)
	case *ast.SelectorExpr:
		v = evalSelector(x, e)
	default:
		panic(fmt.Errorf("cannot eval %T", x))
	}
	return []reflect.Value{v}
}

func eval1(x ast.Expr, e *Env) reflect.Value {
	a := eval(x, e)
	if len(a) != 1 {
		panic("multi- or zero-valued expression in single value context")
	}
	return a[0]
}

func evalBasicLit(x *ast.BasicLit) reflect.Value {
	switch x.Kind {
	case token.STRING:
		s, _ := strconv.Unquote(x.Value)
		return reflect.ValueOf(s)
	default:
		panic(fmt.Errorf("cannot eval %v literal", x.Kind))
	}
}

func evalCall(x *ast.CallExpr, e *Env) []reflect.Value {
	f := eval1(x.Fun, e)
	var in []reflect.Value
	if len(x.Args) == 1 {
		in = append(in, eval(x.Args[0], e)...)
	} else {
		for _, arg := range x.Args {
			in = append(in, eval1(arg, e))
		}
	}
	return f.Call(in)
}

func evalSelector(x *ast.SelectorExpr, e *Env) reflect.Value {
	v := eval1(x.X, e)
	if p, ok := v.Interface().(Package); ok {
		return p.get(x.Sel.Name)
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		return v.FieldByName(x.Sel.Name)
	}
	panic(fmt.Errorf("cannot select from %v", v))
}
