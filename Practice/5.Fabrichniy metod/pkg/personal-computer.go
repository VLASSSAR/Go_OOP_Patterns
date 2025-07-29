package pkg

import "fmt"

type PersonalComputer struct {
	Type    string
	Cpu     int
	Memory  int
	Monitor bool
}

func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Cpu:     20,
		Memory:  64,
		Monitor: true,
	}
}

func (pc PersonalComputer) GetType() string {
	return pc.Type
}

func (pc PersonalComputer) PrintDetails() {
	fmt.Printf("%s, Cpu:[%v], Memory:[%v], Monitor:[%v]\n", pc.Type, pc.Cpu, pc.Memory, pc.Monitor)
}
