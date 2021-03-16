package types

import "reflect"

type BaseOptionValue interface {
	GetOptionValue() *OptionValue
}

func init() {
	t["BaseOptionValue"] = reflect.TypeOf((*OptionValue)(nil)).Elem()
}