package flags

import "strconv"

type intValue int

func (boolean *intValue) Set(value string) error {
	v, err := strconv.ParseInt(value, 10, 0)
	*boolean = intValue(v)
	return err
}

func (boolean *intValue) Get() interface{} {
	return int(*boolean)
}

func Int(short string, long string, defaultValue int, description string) *int {
	pointer := new(int)
	IntVar(pointer, short, long, defaultValue, description)
	return pointer
}

func IntVar(pointer *int, short string, long string, defaultValue int, description string) *int {
	val := (*intValue)(pointer)
	createFlag(val, short, long, strconv.FormatInt(int64(defaultValue), 10), false, description)
	return pointer
}
