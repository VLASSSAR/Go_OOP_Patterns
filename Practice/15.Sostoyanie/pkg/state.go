package pkg

// То какими действиями мы можем взаимодействовать с машиной напрямую.
type State interface {
	RequestItem() error
	InsertMoney(money int) error
	AddItem(count int) error
	DispenseItem() error
}
