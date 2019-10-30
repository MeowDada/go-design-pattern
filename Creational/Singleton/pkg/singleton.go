package singleton

import (
	"sync"
	"fmt"
)

type Singleton struct {
	greeting string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	
	// once.Do(...) will only do once in the program lifetime.
	// So we can do initialization here
	once.Do(func() {
		fmt.Println("Initialize the singleton")
		instance = &Singleton{ greeting:"Hello", }
	})
	return instance
}

func (s Singleton) SayHello() {
	fmt.Println(s.greeting)
}