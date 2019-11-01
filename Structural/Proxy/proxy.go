package main

import "fmt"

type CarInterface interface {
	DriveCar()
} 

type Driver struct {
	age int
}

type Car struct {}
func (c Car) DriveCar() {
	fmt.Println("Car has been driven!")
}

type ProxyCar struct {
	driver Driver
	car    CarInterface
}
func (c ProxyCar) DriveCar() {
	if c.driver.age < 18 {
		fmt.Println("The driver is too young to drive")
	} else {
		c.car.DriveCar()
	}
}

func main() {
	ProxyCar{Driver{16},Car{}}.DriveCar()
	ProxyCar{Driver{25},Car{}}.DriveCar()
}