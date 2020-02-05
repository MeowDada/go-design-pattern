package visitor

type CPU struct { 
	power int
}

func NewCPU(power int) *CPU {
	return &CPU{power}
}

func (c *CPU) Accept(visitor Visitor) {
	visitor.GetCPU(c)
}

func (c *CPU) AssignCPU() int { 
	return c.power
}