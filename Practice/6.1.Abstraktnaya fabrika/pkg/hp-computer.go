package pkg

import "fmt"

type HpComputer struct {
	Cpu    int
	Memory int
}

func (pc HpComputer) PrintDetails() {
	fmt.Printf("[Asus] Cpu:[%d], Memory: [%d]\n", pc.Cpu, pc.Memory)
}
