package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Params struct {
	name, greetLang string
}

func Hello(p Params) string {
	if p.name == "" {
		p.name = "World"
	} else {
		caser := cases.Title(language.English)
		p.name = caser.String(p.name)
	}

	if p.greetLang == "" {
		p.greetLang = "english"
	} else {
		p.greetLang = strings.ToLower(p.greetLang)
	}

	var prefix string
	switch p.greetLang {
	case "french":
		prefix = "Bonjour, "
	case "spanish":
		prefix = "Hola, "
	default:
		prefix = "Hello, "
	}

	return prefix + p.name
}

func main() {
	fmt.Println(Hello(Params{name: "Abhinav", greetLang: "spanish"}))
}
