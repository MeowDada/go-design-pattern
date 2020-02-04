package composite

import (
	"fmt"
)

type Junior struct {
	boss Component
	name string
}

func NewJunior(name string) *Junior {
	return &Junior { name: name }
}

func NewJuniorComponent(name string) Component {
	return NewJunior(name)
}

func (j *Junior) Name() string {
	return j.name
}

func (j *Junior) Boss() Component {
	return j.boss
}

func (j *Junior) SetBoss(boss Component) {
	j.boss = boss
}

func (j *Junior) AddSubordinate(Component) {}

func (j *Junior) Print() {
	fmt.Printf("Hi, I am a junior engineer %s\n", j.name)
}