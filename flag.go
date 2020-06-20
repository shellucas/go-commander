package flags

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/BlackwonderTF/go-flags/utils"
)

type flag struct {
	short    string
	long     string
	required bool
	dfault   string
	value    string
}

var flags map[string]*flag = make(map[string]*flag)
var args map[string]string = make(map[string]string)

func (f flag) IsValid() bool {
	return f.value != ""
}

func (f flag) String() *string {
	return utils.CreateStringPointer(f.value)
}

func (f flag) Bool() *bool {
	b, err := strconv.ParseBool(*f.String())

	if err != nil {
		return nil
	}

	return utils.CreateBoolPointer(b)
}

func ReadFlag(key string) flag {
	return *flags[key]
}

func ParseFlags() {
	arguments := strings.Split(strings.Join(os.Args, " "), "-")

	for i := 0; i < len(arguments); i++ {
		regex, err := regexp.Compile("[a-zA-Z]+ .+")
		if err != nil {
			log.Fatal("Regex could not be compiled")
		}

		if !regex.MatchString(arguments[i]) {
			copy(arguments[i:], arguments[i+1:])
			arguments[len(arguments)-1] = ""
			arguments = arguments[:len(arguments)-1]
			i--
		}
	}

	for _, arg := range arguments {
		pair := strings.Split(strings.TrimSpace(arg), " ")
		args[pair[0]] = pair[1]

		if flags[pair[0]] != nil {
			flags[pair[0]].value = pair[1]
		}
	}
}
