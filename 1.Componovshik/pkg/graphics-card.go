package pkg

import "fmt"

type GraphicsCard struct {
	Name        string
	Description string
}

func (gc GraphicsCard) GetName() string {
	return gc.Name
}

func (gc GraphicsCard) Search(name string) {
	if gc.Name == name {
		fmt.Printf("Компонент [%s] найден %s", gc.Name, gc.Description)
	}
}
