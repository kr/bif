package interp

import (
	"go/parser"
	"reflect"
	"testing"
)

func TestEval1(t *testing.T) {
	type m int
	var e *Env
	e = e.With("s", reflect.ValueOf("s"))
	e = e.With("p", reflect.ValueOf(Package{"X": reflect.ValueOf(1)}))
	e = e.With("t", reflect.ValueOf(struct{ X int }{1}))
	e = e.With("identity", reflect.ValueOf(func(s string) string { return s }))
	e = e.With("m1", reflect.ValueOf(m(1)))
	cases := []struct {
		x string
		w interface{}
	}{
		{`"a"`, "a"},
		{`1`, 1},
		{`s`, "s"},
		{`p.X`, 1},
		{`t.X`, 1},
		{`identity("a")`, "a"},
		{`"a" + "b"`, "ab"},
		{`1 + 1`, 2},
		{`1 - 1`, 0},
		{`m1 + m1`, m(2)},
	}

	for _, test := range cases {
		x, _ := parser.ParseExpr(test.x)
		rv, err := Eval1(x, e)
		if err != nil {
			t.Errorf("Eval(%#q) err: %v", test.x, err)
		}
		if !rv.IsValid() {
			t.Errorf("Eval(%#q) = zero reflect.Value", test.x)
			continue
		}
		if !rv.CanInterface() {
			t.Errorf("Eval(%#q) cannot get interface{}", test.x)
			continue
		}
		g := rv.Interface()
		if !reflect.DeepEqual(g, test.w) {
			t.Errorf("Eval(%#q) = %T %#v want %T %#v", test.x, g, g, test.w, test.w)
		}
	}
}
