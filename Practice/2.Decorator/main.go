package main

import (
	"2.Decorator/pkg"
	"fmt"
)

var (
	base = pkg.BasePc{}

	home = pkg.HomePc{
		Cpu:          4,
		GraphicsCard: 1,
		Wrapper:      base,
	}

	server = pkg.ServerPc{
		Cpu:     16,
		Memory:  256,
		Wrapper: base,
	}
)

func main() {
	fmt.Printf("Base Price: [%f]\n", base.GetPrice())
	fmt.Printf("Home Price: [%f], Cpu:[%d], GraphicsCard:[%d]\n", home.GetPrice(), home.Cpu, home.GraphicsCard)
	fmt.Printf("Server Price: [%f], Cpu:[%d], Memory:[%d]\n", server.GetPrice(), server.Cpu, server.Memory)
}
