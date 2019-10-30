package main

import (
	"fmt"
	
	composite "github.com/MeowDada/go-design-pattern/Structural/Composite/pkg"
)

func main() {

	composite1 := composite.Composite{}
	composite2 := composite.Composite{}

	skill := composite.Skill{}
	exp := composite.Experience{}
	contest := composite.Contest{}

	composite1.AddResume(skill)
	composite1.AddResume(exp)
	composite1.AddResume(contest)

	composite2.AddResume(exp)

	fmt.Println(composite1.DisplayEntry())
	fmt.Println(composite2.DisplayEntry())
}