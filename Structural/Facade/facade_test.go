package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {

	// Create a facade game booster.
	booster := &GameBooster {
		&IntelAPI{},
		&NvidiaAPI{},
		&AsusMotherboardAPI{},
	}

	// Clients only needs to invoke Boost() from the facade game booster.
	// They don't have to worry about the detail of the Boost() method.
	booster.Boost()

}