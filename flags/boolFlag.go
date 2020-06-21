package flags

import "strconv"

type boolValue bool

func (boolean *boolValue) Set(value string) error {
	v, err := strconv.ParseBool(value)
	*boolean = boolValue(v)
	return err
}

func (boolean *boolValue) Get() interface{} {
	return bool(*boolean)
}

func (boolean *boolValue) IsBool() bool {
	return true
}

func Bool(short string, long string, defaultValue bool, description string) *bool {
	pointer := new(bool)
	BoolVar(pointer, short, long, defaultValue, description)
	return pointer
}

func BoolVar(pointer *bool, short string, long string, defaultValue bool, description string) *bool {
	val := (*boolValue)(pointer)
	createFlag(val, short, long, strconv.FormatBool(defaultValue), false, description)
	return pointer
}
