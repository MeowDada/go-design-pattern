package visitor

type Resource interface {
	Accept(Visitor)
}

type Visitor interface {
	GetCPU(*CPU)
	GetDatabase(*Database)
	GetGPU(*GPU)
}

type Admin struct {
	cpuPower int
	databasePower int
	gpuPower int
}

func (a *Admin) GetCPU(cpu *CPU) { 
	a.cpuPower += cpu.AssignCPU() 
}

func (a *Admin) GetDatabase(db *Database) { 
	a.databasePower += db.AssignDatabase() 
}

func (a *Admin) GetGPU(gpu *GPU) { 
	a.gpuPower += gpu.AssignGPU() 
}

func (a *Admin) TotalPower() int { 
	return a.cpuPower + a.databasePower + a.gpuPower 
}

type ResourceManager struct {
	resources []Resource
}

func (rm *ResourceManager) Add(resource Resource) {
	if resource != nil {
		rm.resources = append(rm.resources, resource)
	}
}

func (rm *ResourceManager) Accept(visitor Visitor) {
	for _, resource := range rm.resources {
		resource.Accept(visitor)
	}
}

