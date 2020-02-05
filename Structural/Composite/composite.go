package composite

// Component represents a member in a company. A component is a node in
// an organization tree.
type Component interface {
	Name() string
	Boss() Component
	SetBoss(Component)
	AddSubordinate(Component)
	Print()
}

// Class defines the level of an employee
type Class int

const (
	Highest Class = iota
	Medium
	Lowest
)

// ComponentFactory defines a function to create Component.
type ComponentFactory func(string)Component

var componentFactories = map[Class]ComponentFactory {
	Highest: NewManagerComponent,
	Medium:  NewSeniorComponent,
	Lowest:  NewJuniorComponent,
}

// NewComponent is a factory function to create Component.
func NewComponent(class Class, name string) Component {
	return componentFactories[class](name)
}
