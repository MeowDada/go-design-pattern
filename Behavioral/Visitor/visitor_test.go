package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {

	intelI7 := NewCPU(100)
	intelI3 := NewCPU(40)
	
	gtx2080ti := NewGPU(100)
	gtx1060 := NewGPU(30)

	mysql := NewDatabase(100)
	sqlite := NewDatabase(80)

	rm := &ResourceManager{}
	rm.Add(intelI3)
	rm.Add(gtx1060)
	rm.Add(sqlite)

	admin := &Admin{}
	admin2 := &Admin{}

	rm.Accept(admin)
	t.Log(admin.TotalPower())

	rm.Add(intelI7)
	rm.Add(gtx2080ti)
	rm.Add(mysql)
	
	rm.Accept(admin2)
	t.Log(admin2.TotalPower())
}