package main

import "5.Fabrichniy_metod/pkg"

func main() {
	types := []string{pkg.ServerType, pkg.NotebookType, pkg.PersonalComputerType, "monoblock"}

	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}

}
