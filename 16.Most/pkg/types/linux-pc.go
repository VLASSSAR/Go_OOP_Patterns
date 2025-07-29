package types

import (
	"16.Most/pkg/base"
	"fmt"
)

type LinuxPc struct {
	scanner base.Scanner
}

func (pc *LinuxPc) Scan() {
	fmt.Println("Scan pc linux system")
	pc.scanner.ScanFile()
}

func (pc *LinuxPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
