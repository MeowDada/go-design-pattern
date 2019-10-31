package pc

import "fmt"

type CPU struct {}

func (c *CPU) Overclock() {
	fmt.Println("CPU is overclocking")
}