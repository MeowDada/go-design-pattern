package adapterDemo

import "fmt"

type IPhone struct {
	Connector bool
}

func (p *IPhone) Recharge() {
	if p.Connector {
		fmt.Println("Recharged start")
		fmt.Println("Recharged finished")
	} else {
		fmt.Println("Connect lightning first")
	}
}

func (p *IPhone) UseLightning() {
	p.Connector = true
	fmt.Println("Lightning connected")
}