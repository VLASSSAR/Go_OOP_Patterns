package pkg

import "fmt"

type Notebook struct {
	Type    string
	Cpu     int
	Memory  int
	Monitor bool
}

func NewNotebook() Computer {
	return Notebook{
		Type:   NotebookType,
		Cpu:    8,
		Memory: 16,
	}
}

func (pc Notebook) GetType() string {
	return pc.Type
}

func (pc Notebook) PrintDetails() {
	fmt.Printf("%s, Cpu:[%v], Memory:[%v], Monitor:[%v]\n", pc.Type, pc.Cpu, pc.Memory, pc.Monitor)
}
