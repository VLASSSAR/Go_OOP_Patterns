package main

import (
	"4.Fasad/pkg"
	"fmt"
)

var (
	bank = pkg.Bank{
		Name:  "БАНК",
		Cards: []pkg.Card{},
	}
	card1 = pkg.Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = pkg.Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}
	user = pkg.User{
		Name: "Покупатель-1",
		Card: &card1,
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
		Name: "SHOP",
		Products: []pkg.Product{
			prod,
		},
	}
)

func main() {
	println("[Банк] Выпуск карты")
	bank.Cards = append(bank.Cards, card1, card2)
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
	//Короче, фасадом у нас является функция Sell, так как в ней происходит основное взаимодействие со всеми остальными сервисами
	//Объект shop и функция sell является фасадом над сложной бизнес-логикой по покупке и оплате по безналичному рассчету.
}
