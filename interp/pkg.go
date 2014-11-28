package interp

import (
	"reflect"
)

// TODO(kr): handle types, funcs, etc, not just vars.
type Package map[string]reflect.Value

func (p Package) get(name string) reflect.Value {
	return p[name]
}
