package bridge

import (
	"fmt"
)

// GreetingMethod abstracts the concept of greeting.
type GreetingMethod interface {
	Greet()
}

// GreetingLanguage seperates the implementation from the Greeting.
type GreetingLanguage interface {
	GreetWord() string
}

// English implements the GreetingLanguage interface
type English struct {}

func NewEnglish() *English { return &English{} }

func (e *English) GreetWord() string {
	return "Hello"
}

// Chinese implements the GreetingLanguage interface
type Chinese struct {}

func NewChinese() *Chinese { return &Chinese{} }

func (c *Chinese) GreetWord() string {
	return "你好"
}

// ViaEmail implements the GreethingMethod interface
type ViaEmail struct {
	GreetingLanguage
}

// NewViaEmail is a factory method to create ViaEmail object by
// injecting an inner implementor which implements GreetingLanguage interface
func NewViaEmail(language GreetingLanguage) *ViaEmail {
	return &ViaEmail {
		GreetingLanguage: language,
	}
}

func (ve *ViaEmail) Greet() {
	fmt.Printf("%s via email", ve.GreetWord())
}

// ViaTalk implements the GreethingMethod interface
type ViaTalk struct {
	GreetingLanguage
}


// NewViaTalk is a factory method to create ViaTalk object by
// injecting an inner implementor which implements GreetingLanguage interface
func NewViaTalk(language GreetingLanguage) *ViaTalk {
	return &ViaTalk {
		GreetingLanguage: language,
	}
}

func (vt *ViaTalk) Greet() {
	fmt.Printf("%s via talking", vt.GreetWord())
}