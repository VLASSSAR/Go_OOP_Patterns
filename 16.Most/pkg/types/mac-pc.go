package types

import (
	"16.Most/pkg/base"
	"fmt"
)

type MacPc struct {
	scanner base.Scanner
}

func (pc *MacPc) Scan() {
	fmt.Println("Scan pc Mac system")
	pc.scanner.ScanFile()
}

func (pc *MacPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
