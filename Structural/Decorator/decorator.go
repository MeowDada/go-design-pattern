package main

import (
	"fmt"
	dc "github.com/MeowDada/go-design-pattern/Structural/Decorator/pkg"
)

func main() {
	
	sword := dc.NewSword("katana",15)

	fireSword := dc.NewFireEnchanter(&sword, 5)
	iceFireSword := dc.NewIceEnchanter(&fireSword, 5)

	fmt.Println("sword attack...")
	sword.Attack()

	fmt.Println("fire sword attack...")
	fireSword.Attack()

	fmt.Println("iceFire sword attack...")
	iceFireSword.Attack()
}