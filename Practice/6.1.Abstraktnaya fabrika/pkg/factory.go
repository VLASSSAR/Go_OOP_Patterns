package pkg

import (
	"errors"
	"fmt"
)

// Factory - это типа бренд и наша задача сначала получить бренд, а только потом монитор и/или компьютер этого бренда
type Factory interface {
	GetComputer() Computer
	GetMonitor() Monitor
}

// 1ая функция в main - получаем фабрику
func GetFactory(brand string) (Factory, error) {
	switch brand {
	default:
		return nil, errors.New(fmt.Sprintf("%s бренд - не найден!\n", brand))
	case Asus:
		return &AsusFactory{}, nil
	case Hp:
		return &HpFactory{}, nil
	}

}
