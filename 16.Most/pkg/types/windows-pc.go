package types

import (
	"16.Most/pkg/base"
	"fmt"
)

type WinPc struct {
	scanner base.Scanner
}

func (pc *WinPc) Scan() {
	fmt.Println("Scan pc windows system")
	pc.scanner.ScanFile()
}

func (pc *WinPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
