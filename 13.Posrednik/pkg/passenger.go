package pkg

import "fmt"

// Реализация пасажиров
// Объекты, реализующие интерфейс Dispatcher, также реализуют интерфейс Vehicle
type Passenger struct {
	Dispatcher
}

func (g *Passenger) PermitArrival() {
	fmt.Println("Пассажиры: занимайте места!")
}

func (g *Passenger) Go() {
	fmt.Println("Пассажиры: Отправление!")
	g.Dispatcher.NotifyAboutGo()
}

func (g *Passenger) Arrive() {
	if !g.CanArrive(g) {
		fmt.Println("Пассажиры: отправление задерживается...платформа занята...")
		return
	}
	fmt.Println("Пассажиры: Занимайте места...")
}
