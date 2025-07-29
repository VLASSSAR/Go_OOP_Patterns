package pkg

import "fmt"

type Pc struct {
	Name        string
	Description string
	Components  []Component
}

func (pc Pc) GetName() string {
	return pc.Name
}

func (pc Pc) Search(name string) {
	if pc.Name == name {
		fmt.Printf("Компонент [%s] найден %s", pc.Name, pc.Description)
		return
	}
	for _, component := range pc.Components {
		component.Search(name)
	}
}
