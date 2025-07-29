package pkg

import "fmt"

type AsusComputer struct {
	Cpu    int
	Memory int
}

func (pc AsusComputer) PrintDetails() {
	fmt.Printf("[Asus] Cpu:[%d], Memory: [%d]\n", pc.Cpu, pc.Memory)
}
