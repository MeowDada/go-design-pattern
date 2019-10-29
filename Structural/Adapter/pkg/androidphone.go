package adapterDemo

import "fmt"

type AndroidPhone struct {
	Connector bool
}

func (p *AndroidPhone) Recharge() {
	if p.Connector {
		fmt.Println("Recharge start")
		fmt.Println("Recharge finished")
	} else {
		fmt.Println("Connect MircorUSB first")
	}
}

func (p *AndroidPhone) UseMircoUSB() {
	p.Connector = true
	fmt.Println("MicroUSB connected")
}