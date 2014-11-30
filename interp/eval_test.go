package interp

import (
	"go/parser"
	"reflect"
	"testing"
)

func TestEval1(t *testing.T) {
	e := Env0()
	e = e.Augment("s", reflect.ValueOf("s"))
	e = e.Augment("p", reflect.ValueOf(Package{"X": reflect.ValueOf(1)}))
	e = e.Augment("identity", reflect.ValueOf(func(s string) string { return s }))
	cases := []struct {
		x string
		w interface{}
	}{
		{`"a"`, "a"},
		{`s`, "s"},
		{`p.X`, 1},
		{`identity("a")`, "a"},
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
			t.Errorf("Eval(%#q) = %#v want %#v", test.x, g, test.w)
		}
	}
}
