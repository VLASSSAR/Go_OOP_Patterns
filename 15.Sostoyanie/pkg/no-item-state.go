package pkg

import "fmt"

// Состояние инкапсулирует вендинговую машину, чтобы мы могли с ней непосредственно работать (то есть это то состояние, когда у нас не присутствует товара в машине).
type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) AddItem(count int) error {
	//fmt.Printf("%d items added\n", count)
	i.vendingMachine.IncrementItemCount(count)
	i.vendingMachine.SetState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) DispenseItem() error { //выдача товара
	return fmt.Errorf("Item out of stock")
}
