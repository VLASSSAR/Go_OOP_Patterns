package main

import (
	"6.Abstraktnaya_fabrika/pkg"
)

var brands = []string{pkg.Asus, pkg.Hp, "dell"}

func main() {

	for _, brand := range brands {
		factory, err := pkg.GetFactory(brand)
		if err != nil {
			println(err.Error())
			continue
		}
		computer := factory.GetComputer()
		computer.PrintDetails()

		monitor := factory.GetMonitor()
		monitor.PrintDetails()
	}
}
