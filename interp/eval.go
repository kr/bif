package interp

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)

// A nil *Env represents the empty environment.
type Env struct {
	s string
	v reflect.Value
	e *Env
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

func (e *Env) With(name string, v reflect.Value) *Env {
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
	// TODO(kr): untyped constant expressions
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
	case *ast.BinaryExpr:
		v = evalBinary(x, e)
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
	case token.INT:
		// TODO(kr): treat untyped constant values correctly.
		// for now we convert an INT literal to an int.
		n, _ := strconv.Atoi(x.Value)
		return reflect.ValueOf(n)
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
		// TODO(kr): unexported fields in target package
		return v.FieldByName(x.Sel.Name)
	}
	panic(fmt.Errorf("cannot select from %v", v))
}

func evalBinary(exp *ast.BinaryExpr, env *Env) reflect.Value {
	x := eval1(exp.X, env)
	y := eval1(exp.Y, env)
	if x.Type() != y.Type() {
		panic(fmt.Errorf("mismatched types %v and %v", x.Type(), y.Type()))
	}
	f, ok := binops[exp.Op][x.Kind()]
	if !ok {
		panic(fmt.Errorf("operator %v not defined on %v", exp.Op, x.Kind()))
	}
	x0 := x.Convert(kindtypes[x.Kind()])
	y0 := y.Convert(kindtypes[y.Kind()])
	v := f(x0.Interface(), y0.Interface())
	return reflect.ValueOf(v).Convert(x.Type())
}
