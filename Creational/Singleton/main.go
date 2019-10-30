package main

import (
	singleton "github.com/MeowDada/go-design-pattern/Creational/Singleton/pkg"
)

func main() {

	instance := singleton.GetInstance()
	instance.SayHello()
	singleton.GetInstance().SayHello()

}