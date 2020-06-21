package flags

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/BlackwonderTF/go-commander/utils"
)

type flag struct {
	value string
}

var version string
var flags map[string]*flag = parseFlags()

func checkRequired(f *flag, boolean bool) {
	if f != nil && (boolean || f.value != "") {
		return
	}

	usageString := "placeholder"
	log.Fatal(usageString)
}

func getFlag(short string, long string) *flag {
	var f *flag

	fs := flags[short]
	fl := flags[long]

	if fs != nil {
		f = fs
	} else if fl != nil {
		f = fl
	} else {
		f = nil
	}

	return f
}

// String defines a flag
func String(short string, long string, baseValue string, description string) string {
	f := getFlag(short, long)
	if f == nil || f.value == "" {
		return baseValue
	}
	return f.value
}

// StringRequired defines a flag that is required
func StringRequired(short string, long string, description string) string {
	f := getFlag(short, long)
	checkRequired(f, false)
	return f.value
}

// Bool defines a boolean switch flag
func Bool(short string, long string, description string) bool {
	f := getFlag(short, long)
	return f != nil
}

// parseFlags parses the command line arguments.
func parseFlags() map[string]*flag {
	argumentsRegex := regexp.MustCompile("(-[a-zA-Z]+(\\s[^-\\s]+){0,1}|--[a-zA-Z]+(=[^\\s]+){0,1})")
	arguments := argumentsRegex.FindAllString(strings.Join(os.Args[1:], " "), -1)

	args := make(map[string]*flag)

	for _, arg := range arguments {
		pair := utils.RegSplit(strings.TrimSpace(arg), "[\\s=]")
		var pairs []string

		r := regexp.MustCompile("^-[a-zA-Z]+")
		if r.MatchString(pair[0]) {
			pair[0] = strings.ReplaceAll(pair[0], "-", "")
			for _, key := range pair[0] {
				pairs = append(pairs, string(key))
			}
		} else {
			pair[0] = strings.ReplaceAll(pair[0], "-", "")
			pairs = append(pairs, pair[0])
		}

		var value string
		if len(pair) == 2 {
			value = pair[1]
		} else {
			value = ""
		}

		for _, key := range pairs {
			args[key] = new(flag)
			args[key].value = value
		}
	}

	return args
}
