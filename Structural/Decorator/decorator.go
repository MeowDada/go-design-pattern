package decorator

// Beverage is a interface that define a beverage.
type Beverage interface {
	Cost() int
	Description() string
}

// DarkRoast is a concret instance of a Beverage.
// It could be decorated by any numbers of decorators.
type DarkRoast struct {}
func (d *DarkRoast) Cost() int { return 10 }
func (d *DarkRoast) Description() string { return "Dark Roast" }

// Expresso is a concret instance of a Beverage.
// It could be decorated by any numbers of decorators.
type Expresso struct {}
func (e *Expresso) Cost() int { return 15 }
func (e *Expresso) Description() string { return "Expresso" }

// Condiment is a interface that implements Beverage interface.
// It is also act like a decorator.
type Condiment interface {
	Beverage
}

// Mocha implements Condiment interface.
// It could be used as a instance of decorator.
type Mocha struct {
	beverage Beverage
}
func NewMocha(beverage Beverage) *Mocha { return &Mocha{beverage} }
func (m *Mocha) Cost() int { return m.beverage.Cost() + 25 }
func (m *Mocha) Description() string { return "Mocha " + m.beverage.Description() }

// Milk implements Condiment interface.
// It could be used as a instance of decorator.
type Milk struct {
	beverage Beverage
}
func NewMilk(beverage Beverage) *Milk { return &Milk{beverage} }
func (m *Milk) Cost() int { return m.beverage.Cost() + 5 }
func (m *Milk) Description() string { return "Milk " + m.beverage.Description() }

