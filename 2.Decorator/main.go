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
		Wrapper:      base, //Вот в чем прикол! Передаем в обертку базовую комплектацию. Поэтому метод GetPrice будет возвращать 10.0!
	}
	server = pkg.ServerPc{
		Cpu:     16,
		Memory:  256,
		Wrapper: base,
	}
)

func main() {
	fmt.Printf("Base Price:[%f]\n", base.GetPrice())
	fmt.Printf("Home Cpu:[%d] Cards:[%d] Price:[%f]\n", home.Cpu, home.GraphicsCard, home.GetPrice())
	fmt.Printf("Home Cpu:[%d] Memory:[%d] Price:[%f]\n", server.Cpu, server.Memory, server.GetPrice())
}
