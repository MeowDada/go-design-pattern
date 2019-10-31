package dc

import "fmt"

type FireEnchanter struct {
	weapon Weapon
	atk int
}

func NewFireEnchanter(weapon Weapon, atk int) FireEnchanter {
	return FireEnchanter {
		weapon: weapon,
		atk: atk,
	}
}

func (fe *FireEnchanter) Attack() {
	fe.weapon.Attack()
	fmt.Printf("deal additional %d damages due to the burst of fire\n", fe.atk)
}

func (fe *FireEnchanter) GetHitPoint() int {
	return fe.atk + fe.weapon.GetHitPoint() 
}

func (fe *FireEnchanter) SpecialAttack() {

}