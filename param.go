package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Parameter struct {
	Name     string
	Required bool
	Default  string
	UseStdIn bool
}

func Param(name string, required bool, defaultValue string) Parameter {
	return Parameter{
		Name:     name,
		Required: required,
		Default:  defaultValue,
		UseStdIn: false,
	}
}

func StdInParam(name string, required bool) Parameter {
	return Parameter{
		Name:     name,
		Required: required,
		Default:  "",
		UseStdIn: true,
	}
}

type Params struct {
	parameters []Parameter
	values     map[string]string
}

func (p Params) Get(name string) string {
	return p.values[name]
}

func ParseArgs(args []string, parameters []Parameter) (Params, error) {
	values := make(map[string]string)

	for i := 0; i < len(parameters); i++ {
		param := parameters[i]
		var value string

		if i < len(args) {
      value = args[i]
		} else {
      value = ""
		}

		if value == "" && param.Required {
			if param.UseStdIn {
				value = readStdIn()
			} else if param.Default != "" {
				value = param.Default
			} else {
				return Params{}, fmt.Errorf("%s is required", param.Name)
			}
		}

		values[param.Name] = value
	}

	return Params{
		parameters: parameters,
		values:     values,
	}, nil
}

func readStdIn() string {
	value := strings.Builder{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value.WriteString(scanner.Text())
	}

	return value.String()
}
