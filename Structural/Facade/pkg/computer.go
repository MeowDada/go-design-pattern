package pc

type Computer struct {
	cpu CPU
	mem Memory
	gc  GraphicCard
}

func (c *Computer) EnableGameMode() {
	c.cpu.Overclock()
	c.mem.Boost()
	c.gc.Render()
}