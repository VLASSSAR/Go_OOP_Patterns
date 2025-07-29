package main

import (
	"4.Fasad/pkg"
	"fmt"
)

var (
	bank = pkg.Bank{
		Name:  "Tinkoff",
		Cards: []pkg.Card{},
	}
	card = pkg.Card{
		Name:    "CRD-1",
		Balance: 150,
		Bank:    &bank,
	}
	card2 = pkg.Card{
		Name:    "CRD-2",
		Balance: 10,
		Bank:    &bank,
	}
	user = pkg.User{
		Name: "Покупатель-1",
		Card: &card,
	}
	user2 = pkg.User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = pkg.Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = pkg.Shop{
		Name: "Пятерочка",
		Products: []pkg.Product{
			prod,
		},
	}
)

func main() {
	fmt.Printf("[Банк] Выпуск карты!\n")
	bank.Cards = append(bank.Cards, card, card2)
	fmt.Printf("[%s]\n", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
