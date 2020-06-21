package utils

func CreateStringPointer(s string) *string {
	value := new(string)
	*value = s
	return value
}

func CreateBoolPointer(b bool) *bool {
	value := new(bool)
	*value = b
	return value
}
