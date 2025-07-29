package main

import (
	"7.Builder/pkg"
	"fmt"
)

func main() {
	asusCollector := pkg.GetCollector("asus")
	//hpCollector := pkg.GetCollector("hp")

	factory := pkg.NewFactory(asusCollector)

	fmt.Printf("%v", factory)

	//asusComputer := factory.CreateComputer()
	//asusComputer.Print()
	//
	//
	//factory.SetCollector(hpCollector)
	//hpComputer := factory.CreateComputer()
	//hpComputer.Print()
}
