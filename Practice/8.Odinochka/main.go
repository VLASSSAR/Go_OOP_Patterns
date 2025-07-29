package main

import (
	"8.Odinochka/pkg"
	"fmt"
)

var singleton *pkg.Singleton

// Функция, которая вызывается перед main (статический конструктор)
func init() {
	println("init - базовая инициализация программы")
	singleton = &pkg.Singleton{Type: "Одиночка"}
}

func main() {
	for i := 0; i < 3; i++ {
		pkg.NewSingleton(singleton, "Create singleton")
	}
	fmt.Println(singleton)
}

//Это был пример паттерна одиночка, когда мы можем создать только один экземпляр объекта!!!
