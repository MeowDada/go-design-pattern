package composite

import (
	"fmt"
)

type Manager struct {
	name string
	subordinates []Component
}

func NewManager(name string) *Manager {
	return &Manager { name: name }
}

func NewManagerComponent(name string) Component {
	return NewManager(name)
}

func (m *Manager) Name() string {
	return m.name
}

func (m *Manager) Boss() Component {
	return nil
}

func (m *Manager) SetBoss(boss Component) {}

func (m *Manager) AddSubordinate(subordinate Component) {
	m.subordinates = append(m.subordinates, subordinate)
}

func (m *Manager) Print() {
	fmt.Printf("Hi, I am a manager %s\n", m.name)
	fmt.Printf("My subordinates are:\n")
	for _, subordinate := range m.subordinates {
		subordinate.Print()
	}
}