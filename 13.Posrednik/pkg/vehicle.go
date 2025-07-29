package pkg

type Vehicle interface {
	Arrive()
	Go()
	PermitArrival()
}

type Dispatcher interface {
	CanArrive(Vehicle) bool
	NotifyAboutGo()
}
