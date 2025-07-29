package pkg

import (
	"fmt"
	"time"
)

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (card Card) CheckBalance() error {
	fmt.Printf("[Карта] Запрос в банк для проверки баланса\n")
	time.Sleep(time.Millisecond * 300)
	return card.Bank.CheckBalance(card.Name)
}
