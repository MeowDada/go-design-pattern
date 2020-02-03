package bridge

import (
	"fmt"
)

type GreetingMethod interface {
	Greet()
}

type GreetingLanguage interface {
	GreetWord() string
}

type English struct {}

func NewEnglish() *English { return &English{} }

func (e *English) GreetWord() string {
	return "Hello"
}

type Chinese struct {}

func NewChinese() *Chinese { return &Chinese{} }

func (c *Chinese) GreetWord() string {
	return "你好"
}

type ViaEmail struct {
	GreetingLanguage
}

func NewViaEmail(language GreetingLanguage) *ViaEmail {
	return &ViaEmail {
		GreetingLanguage: language,
	}
}

func (ve *ViaEmail) Greet() {
	fmt.Printf("%s via email", ve.GreetWord())
}

type ViaTalk struct {
	GreetingLanguage
}

func NewViaTalk(language GreetingLanguage) *ViaTalk {
	return &ViaTalk {
		GreetingLanguage: language,
	}
}

func (vt *ViaTalk) Greet() {
	fmt.Printf("%s via talking", vt.GreetWord())
}