package types

import "reflect"

func (b *OptionType) GetOptionType() *OptionType { return b }

type BaseOptionType interface {
	GetOptionType() *OptionType
}

func init() {
	t["BaseOptionType"] = reflect.TypeOf((*OptionType)(nil)).Elem()
}

func (b *OptionValue) GetOptionValue() *OptionValue { return b }

type BaseOptionValue interface {
	GetOptionValue() *OptionValue
}

func init() {
	t["BaseOptionValue"] = reflect.TypeOf((*OptionValue)(nil)).Elem()
}
