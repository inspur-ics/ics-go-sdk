package types

import (
	"reflect"
	"strings"
)

var t = map[string]reflect.Type{}

func Add(name string, kind reflect.Type) {
	t[name] = kind
}

type Func func(string) (reflect.Type, bool)

func TypeFunc() Func {
	return func(name string) (reflect.Type, bool) {
		typ, ok := t[name]
		if !ok {
			name = strings.TrimPrefix(name, "types:")
			typ, ok = t[name]
		}
		return typ, ok
	}
}
