package pkg

// Здесь мы создадим общий интерфейс состояний.

type State interface {
	AddItem(int) error
	RequestItem() error // запрос предмета
	InsertMoney(money int) error
	DispenseItem() error //выдача предмета
}
