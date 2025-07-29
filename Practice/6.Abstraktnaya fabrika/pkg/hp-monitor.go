package pkg

import "fmt"

type HpMonitor struct {
	Size int
}

func (pc HpMonitor) PrintDetails() {
	fmt.Printf("[Hp] Monitor Size:[%d]\n", pc.Size)
}
