package dc

import "fmt"

type Sword struct {
	name string
	atk  int
}

func NewSword(name string, atk int) Sword {
	return Sword {
		name: name,
		atk:  atk,
	}
}

func (s *Sword) Attack() {
	fmt.Printf("%s deal %d damages\n", s.name, s.atk)
}

func (s *Sword) GetHitPoint() int {
	return s.atk
}