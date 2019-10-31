package pc

import "fmt"

type Memory struct {}

func (m *Memory) Boost() {
	fmt.Println("Memory is boosting")
}

