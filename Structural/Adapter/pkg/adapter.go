package adapterDemo

import "fmt"

type LightningToMircoUSBAdapter struct {
	LightningPhone LightningInterface
}

func (a *LightningToMircoUSBAdapter) Recharge() {
	a.LightningPhone.Recharge()
}

func (a *LightningToMircoUSBAdapter) UseMicroUSB() {
	a.LightningPhone.UseLightning()
	fmt.Println("Lightning to MircoUSB adapter connected")
}