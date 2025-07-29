package pkg

import "fmt"

type Cargo struct {
	Dispatcher
}

func (g *Cargo) PermitArrival() {
	fmt.Println("Грузовик: погрузка...")
	g.Arrive()
}

func (g *Cargo) Go() {
	fmt.Println("Грузовик: Успешно погружен!")
	g.Dispatcher.NotifyAboutGo()
}

func (g *Cargo) Arrive() {
	if !g.CanArrive(g) {
		fmt.Println("Грузовик: погрузка запрещена, на платформе пассажиры!")
		return
	}
	fmt.Println("Грузовик: Отправлен!")
}
