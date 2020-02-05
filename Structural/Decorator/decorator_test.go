package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {

	// Meal a cup of Expresso.
	base := new(Expresso)
	t.Log(base.Description(), ",cost:", base.Cost())

	// Add some milk into expresso.
	milk := NewMilk(base)
	t.Log(milk.Description(), ",cost:", milk.Cost())

	// Add some mocha!!!
	mocha := NewMocha(milk)
	t.Log(mocha.Description(), ",cost:", mocha.Cost())
}