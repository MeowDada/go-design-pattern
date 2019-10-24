package main

import (
	"fmt"
)

type Mouse interface {
	Brand() string
}

type AsusMouse struct {}
func (m *AsusMouse) Brand() string {
	return "Asus"
}

type DellMouse struct {}
func (m *DellMouse) Brand() string {
	return "Dell"
}

type MouseFactory interface {
	ProduceMouse() Mouse
}

type AsusMouseFactory struct{}
func (f *AsusMouseFactory) ProduceMouse() Mouse {
	return new(AsusMouse)
}

type DellMouseFactory struct{}
func (f *DellMouseFactory) ProduceMouse() Mouse {
	return new(DellMouse)
}

func main() {

	var mouseFactory MouseFactory
	mouseFactory = new(AsusMouseFactory)

	mouse := mouseFactory.ProduceMouse()
	fmt.Println(mouse.Brand())
}