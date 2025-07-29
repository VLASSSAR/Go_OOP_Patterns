package main

import "7.Builder/pkg"

func main() {
	asusCollector := pkg.GetCollector(pkg.AsusCollectorType)
	hpCollector := pkg.GetCollector(pkg.HpCollectorType)

	factory := pkg.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()
}
