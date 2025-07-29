package pkg

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.SetState(i.vendingMachine.hasItem)
	}
	return nil
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) InsertMoney(count int) error {
	return fmt.Errorf("Item out of stock")
}
