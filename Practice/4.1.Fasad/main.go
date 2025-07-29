package main

import (
	"4.Fasad/pkg"
	"fmt"
)

var (
	bank = pkg.Bank{
		Name:  "Тинькоф",
		Cards: []pkg.Card{},
	}
	card = pkg.Card{
		Name:    "Покупатель-1",
		Balance: 150,
		Bank:    &bank,
	}
	card2 = pkg.Card{
		Name:    "Покупатель-2",
		Balance: 5,
		Bank:    &bank,
	}
	user = pkg.User{
		Name: "Мартин",
		Card: &card,
	}
	user2 = pkg.User{
		Name: "Колян",
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
	err := shop.Sell(user, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
