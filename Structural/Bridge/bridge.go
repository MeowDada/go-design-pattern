package main

import (
	"fmt"
	language "github.com/MeowDada/go-design-pattern/Structural/Bridge/pkg"
)

type TaiwanPeople struct {
	GreetingStyle language.Greeting
}

func (tp *TaiwanPeople) SayHello() {
	fmt.Println(tp.GreetingStyle.Hello())
}

func (tp *TaiwanPeople) SayThanks() {
	fmt.Println(tp.GreetingStyle.Thanks())
}

func main() {

	people := TaiwanPeople {
		GreetingStyle: &language.ChineseGreeting{},
	}

	people.SayHello()
	people.SayThanks()
}