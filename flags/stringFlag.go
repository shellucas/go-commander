package flags

type stringValue string

func (str *stringValue) Set(value string) error {
	*str = stringValue(value)
	return nil
}

func (str *stringValue) Get() interface{} {
	return string(*str)
}

func String(short string, long string, defaultValue string, description string) *string {
	pointer := new(string)
	StringVar(pointer, short, long, defaultValue, description)
	return pointer
}

func StringVar(pointer *string, short string, long string, defaultValue string, description string) *string {
	val := (*stringValue)(pointer)
	createFlag(val, short, long, defaultValue, false, description)
	return pointer
}

func StringRequired(short string, long string, description string) *string {
	pointer := new(string)
	StringVarRequired(pointer, short, long, description)
	return pointer
}

func StringVarRequired(pointer *string, short string, long string, description string) *string {
	val := (*stringValue)(pointer)
	createFlag(val, short, long, "", true, description)
	return pointer
}
