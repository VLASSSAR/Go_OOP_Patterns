package main

import (
	"15.Sostoyanie/pkg"
	"fmt"
	"log"
)

func main() {
	vendingMachine := pkg.NewVendingMachine(2, 10)
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem() //return HasItemState
	if err != nil {
		log.Fatalf(err.Error())
	}
	//
	//fmt.Println(vendingMachine.CurrentState)
	//
	fmt.Println()
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem() //return HasItemState
	if err != nil {
		log.Fatalf(err.Error())
	}
}
