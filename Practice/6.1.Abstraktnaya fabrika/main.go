package main

import "6.Abstraktnaya_fabrika/pkg"

var (
	brands = []string{pkg.Asus, pkg.Hp, "dell"}
)

func main() {
	for _, brand := range brands {
		factory, err := pkg.GetFactory(brand) //получили бренд-объект (AsusFactory{} || HpFactory{}) у которого хотим получить данные о мониторе и пк
		if factory == nil {
			println(err.Error())
		}
		computer := factory.GetComputer() //получили ПК у бренда (вернулся объект поддерживающий интерфейс Computer)
		computer.PrintDetails()           //вернули детали этого объекта

		monitor := factory.GetMonitor() //то же самое, только с монитором!
		monitor.PrintDetails()

	}
}
