package main

import (
	adapterDemo "github.com/MeowDada/go-design-pattern/Structural/Adapter/pkg"
)

func main() {

	iphone := adapterDemo.IPhone{Connector:false}
	adapter := adapterDemo.LightningToMircoUSBAdapter{ LightningPhone: &iphone, }
	adapter.UseMicroUSB()
	adapter.Recharge()
}