package composite

import (
	"fmt"
	"testing"
)

func TestComponent(t *testing.T) {

	junior1 := NewComponent(Lowest, "Tim")
	junior2 := NewComponent(Lowest, "Leo")
	junior3 := NewComponent(Lowest, "Peter")
	junior4 := NewComponent(Lowest, "Zack")

	senior1 := NewComponent(Medium, "John")
	senior2 := NewComponent(Medium, "Leon")

	manager := NewComponent(Highest, "Jack")

	junior1.SetBoss(senior1)
	junior2.SetBoss(senior1)
	senior1.AddSubordinate(junior1)
	senior1.AddSubordinate(junior2)
	senior1.SetBoss(manager)

	junior3.SetBoss(senior2)
	junior4.SetBoss(senior2)
	senior2.AddSubordinate(junior3)
	senior2.AddSubordinate(junior4)
	senior2.SetBoss(manager)

	manager.AddSubordinate(senior1)
	manager.AddSubordinate(senior2)

	junior1.Print()
	fmt.Println()
	senior1.Print()
	fmt.Println()
	manager.Print()
}