package main

import "5.Fabrichniy_metod/pkg"

var types = []string{pkg.PersonalComputerType, pkg.NotebookType, pkg.ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}
