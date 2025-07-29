package main

import "Componovshik/pkg"

var (
	cpu = pkg.Cpu{
		Name:        "CP-1",
		Description: "Процессор 1",
	}
	cpu2 = pkg.Cpu{
		Name:        "CP-2",
		Description: "Процессор 2",
	}
	card = pkg.GraphicsCard{
		Name:        "Radeon",
		Description: "Видеокарта - 1",
	}
	card2 = pkg.GraphicsCard{
		Name:        "Geforce",
		Description: "Видеокарта - 2",
	}
	motherboard = pkg.Motherboard{
		Name:        "GigaByte",
		Description: "Материнская плата",
		Components: []pkg.Component{
			cpu, cpu2, card, card2,
		},
	}
	pc = pkg.Pc{
		Name:        "PC",
		Description: "Персональный компьютер",
		Components: []pkg.Component{
			motherboard,
		},
	}
)

func main() {
	pc.Search("Geforce")
}
