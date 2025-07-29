package pkg

import (
	"errors"
	"fmt"
)

type Factory interface {
	GetMonitor() Monitor
	GetComputer() Computer
}

func GetFactory(brand string) (Factory, error) {
	switch brand {
	default:
		return nil, errors.New(fmt.Sprintf("Производитель %s - не найден!", brand))

	case Asus:
		return &AsusFactory{}, nil
	case Hp:
		return &HpFactory{}, nil
	}
}
