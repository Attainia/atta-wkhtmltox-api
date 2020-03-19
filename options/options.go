package options

import (
	"strings"
)

type Options struct {
	options map[string]string
}

func OptionsFromHeader(header string) (*Options, error) {
	headerParts := strings.Split(header, ";")

	options := map[string]string{}
	if len(headerParts) > 1 {
		optionsString := strings.Split(headerParts[1], " ")
		for i := range optionsString[1:] {
			optionParts := strings.Split(optionsString[i+1], "=")
			options[optionParts[0]] = optionParts[1]
		}
	}

	return &Options{options}, nil
}
