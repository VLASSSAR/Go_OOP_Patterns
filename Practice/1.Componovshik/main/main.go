package main

import "1.Componovshik/pkg"

func main() {

	var (
		cpu = pkg.Cpu{
			Name:        "Intel Core I7-12700k",
			Description: "Процессор Intel",
		}
		cpu2 = pkg.Cpu{
			Name:        "AMD Ryzen 5990X",
			Description: "Процессор AMD",
		}
		gc = pkg.GraphicsCard{
			Name:        "RTX 4080",
			Description: "Видеокарта Geforce",
		}
		gc2 = pkg.GraphicsCard{
			Name:        "RTX 4070",
			Description: "Видеокарта Geforce",
		}
		motherboard = pkg.Motherboard{
			Name:        "ASUS",
			Description: "Материнская карта ASUS",
			Components: []pkg.Component{
				cpu, cpu2, gc, gc2,
			},
		}
		pc = pkg.Pc{
			Name:        "Home_PC",
			Description: "Домашняя база",
			Components: []pkg.Component{
				motherboard,
			},
		}
	)

	pc.Search("RTX 4080")

}
