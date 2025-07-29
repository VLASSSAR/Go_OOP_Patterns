package pkg

import "fmt"

type Computer struct {
	Core         int
	Memory       int
	GraphicsCard int
	Monitor      int
	Brand        string
}

func (pc *Computer) Print() {
	fmt.Printf("%s Core:[%d] Mem:[%d] GraphicsCard:[%d] Monitor:[%d]\n", pc.Brand, pc.Core, pc.Memory, pc.Monitor, pc.GraphicsCard)
}
