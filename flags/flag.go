package flags

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/shellucas/go-commander/utils"
)

type value interface {
	Set(string) error
	Get() interface{}
}

type flag struct {
	short        string
	long         string
	required     bool
	description  string
	value        value
	defaultValue string
}

var allFlags []*flag = []*flag{}
var shortFlags map[string]*flag = make(map[string]*flag)
var longFlags map[string]*flag = make(map[string]*flag)
var help *bool

func init() {
	help = Bool("", "help", false, "Shows the usage of the program")
}

func createFlag(value value, short string, long string, defaultValue string, required bool, description string) *flag {
	f := &flag{
		short:        short,
		long:         long,
		required:     required,
		description:  description,
		value:        value,
		defaultValue: defaultValue,
	}

	if shortFlags[short] != nil {
		log.Fatalf("Multiple flags with short flag \"%s\" exist", short)
	}

	if longFlags[long] != nil {
		log.Fatalf("Multiple flags with long flag \"%s\" exist", long)
	}

	if short != "" {
		shortFlags[short] = f
	}

	if long != "" {
		longFlags[long] = f
	}

	allFlags = append(allFlags, f)

	f.value.Set(defaultValue)

	return f
}

func Parse() {
	argumentsRegex := regexp.MustCompile("(-[a-zA-Z]+(\\s[^-\\s]+){0,1}|--[a-zA-Z]+(=[^\\s]+){0,1})")
	arguments := argumentsRegex.FindAllString(strings.Join(os.Args[1:], " "), -1)

	for _, arg := range arguments {
		pair := utils.RegSplit(strings.TrimSpace(arg), "[\\s=]")

		r := regexp.MustCompile("^-[a-zA-Z]+")
		copy := pair[0]
		pair[0] = strings.ReplaceAll(pair[0], "-", "")

		if r.MatchString(copy) {
			counter := 0
			for _, option := range pair[0] {
				f := shortFlags[string(option)]

				if f == nil {
					continue
				}

				if val, hasBoolFlag := f.value.(*boolValue); hasBoolFlag && val.IsBool() {
					f.value.Set("true")
				} else {
					f.value.Set(pair[1])
					counter++
				}

				if counter > 1 {
					// TODO make a better output error message and/or system
					log.Fatal("Multiple inputs with same value given, is not allowed")
				}
			}
		} else {
			f := longFlags[pair[0]]

			if f == nil {
				continue
			}

			if len(pair) == 1 {
				f.value.Set("true")
			} else {
				f.value.Set(pair[1])
			}
		}
	}

	if help != nil && *help {
		fmt.Println(createUsage())
		os.Exit(0)
	}

	for _, flag := range allFlags {
		if flag.required && flag.value.Get() == "" {
			fmt.Println(createUsage())
			os.Exit(0)
		}
	}
}

func createUsage() string {
	builder := strings.Builder{}

	// TODO make this pretty pl0x
	for _, f := range allFlags {
		if f.short != "" {
			builder.WriteString(fmt.Sprintf("\n\n-%s \t %s", f.short, f.description))
		}

		if f.long != "" {
			builder.WriteString(fmt.Sprintf("\n\n--%s \t %s", f.long, f.description))
		}
	}

	return builder.String()
}
