package main

import (
	builder "github.com/MeowDada/go-design-pattern/Creational/Builder/pkg"
	"fmt"
)

func main() {

	bikeBuilder := builder.BikeBuilder{}
	
	giant := new(builder.ManufacturingDirector)
	giant.SetBuilder(&bikeBuilder)
	giant.Construct()
	fmt.Println(giant.GetProduct())

	carBuilder := builder.CarBuilder{}

	nissan := new(builder.ManufacturingDirector)
	nissan.SetBuilder(&carBuilder)
	nissan.Construct()
	fmt.Println(nissan.GetProduct())
}