package dc

import "fmt"

type IceEnchanter struct {
	weapon Weapon
	atk    int
}

func NewIceEnchanter(weapon Weapon, atk int) IceEnchanter {
	return IceEnchanter {
		weapon: weapon,
		atk: atk,
	}
}


func (ie *IceEnchanter) Attack() {
	ie.weapon.Attack()
	fmt.Printf("deal additional %d damages due to the burst of ice\n", ie.atk)
}

func (ie *IceEnchanter) GetHitPoint() int {
	return ie.atk + ie.weapon.GetHitPoint() 
}

func (ie *IceEnchanter) SpecialAttack() {

}