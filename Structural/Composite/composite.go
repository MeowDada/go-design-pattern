package composite

type Component interface {
	Name() string
	Boss() Component
	SetBoss(Component)
	AddSubordinate(Component)
	Print()
}

type Class int

const (
	Highest Class = iota
	Medium
	Lowest
)

type ComponentFactory func(string)Component

var componentFactories = map[Class]ComponentFactory {
	Highest: NewManagerComponent,
	Medium:  NewSeniorComponent,
	Lowest:  NewJuniorComponent,
}

func NewComponent(class Class, name string) Component {
	return componentFactories[class](name)
}
