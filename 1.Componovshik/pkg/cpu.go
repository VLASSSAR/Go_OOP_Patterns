package pkg

import "fmt"

type Cpu struct {
	Name        string
	Description string
}

func (cpu Cpu) GetName() string {
	return cpu.Name
}

func (cpu Cpu) Search(name string) {
	if cpu.Name == name {
		fmt.Printf("Компонент [%s] найден %s", cpu.Name, cpu.Description)
	}
}
