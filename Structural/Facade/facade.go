package facade

import (
	"fmt"
)

type GameBooster struct {
	CentralProcessUnitAPI
	GraphicAPI
	MotherboardAPI
}

func (g *GameBooster) Boost() {
	g.CPUBoost()
	g.GPUBoost()
	g.FanBoost()
}

type CentralProcessUnitAPI interface { CPUBoost() }
type IntelAPI struct {}
func (i IntelAPI) CPUBoost() { fmt.Println("CPU is boosting") }

type GraphicAPI interface { GPUBoost() }
type NvidiaAPI struct {}
func (n NvidiaAPI) GPUBoost() { fmt.Println("GPU is boosting") }

type MotherboardAPI interface { FanBoost() }
type AsusMotherboardAPI struct {}
func (a AsusMotherboardAPI) FanBoost() { fmt.Println("Fans are boosting") }

