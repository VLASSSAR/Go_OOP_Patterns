package pkg

import "fmt"

// ПК - родитель для всех остальных компонентов (для материнской платы, для процессора, для видеокарты)
type Pc struct {
	Name        string
	Description string
	Components  []Component
}

// Также как и у материнской платы, мы должны реализовать все необходимые методы и функции базового интерфейса компонента.

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
