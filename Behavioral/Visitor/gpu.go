package visitor

type GPU struct { 
	power int 
}

func NewGPU(power int) *GPU { 
	return &GPU{power} 
}

func (g *GPU) Accept(visitor Visitor) { 
	visitor.GetGPU(g) 
}

func (g *GPU) AssignGPU() int { 
	return g.power 
}