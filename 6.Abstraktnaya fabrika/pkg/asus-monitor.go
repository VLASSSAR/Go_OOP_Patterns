package pkg

import "fmt"

type AsusMonitor struct {
	Size int
}

func (pc AsusMonitor) PrintDetails() {
	fmt.Printf("[Asus] Monitor Size:[%d]\n", pc.Size)
}
