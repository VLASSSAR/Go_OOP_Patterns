package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю для получения остатка")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance() //У карты пользователя проверяем баланс
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка - может ли [%s] купить товар", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара")
		}
		fmt.Printf("[Магазин] Товар [%s] - выкуплен\n", prod.Name)
	}
	return nil
}
