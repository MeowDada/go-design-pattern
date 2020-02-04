package composite

import (
	"fmt"
)

type Senior struct {
	boss Component
	name string
	subordinates []Component
}

func NewSenior(name string) *Senior {
	return &Senior { name: name }
}

func NewSeniorComponent(name string) Component {
	return NewSenior(name)
}

func (s *Senior) Name() string {
	return s.name
}

func (s *Senior) Boss() Component {
	return s.boss
}

func (s *Senior) SetBoss(boss Component) {
	s.boss = boss
}

func (s *Senior) AddSubordinate(subordinate Component) {
	s.subordinates = append(s.subordinates, subordinate)
}

func (s *Senior) Print() {
	fmt.Printf("Hi, I am a senior engineer %s\n", s.name)
	fmt.Printf("My subordinates are:\n")
	for _, subordinate := range s.subordinates {
		subordinate.Print()
	}
}