package main

import "9.Cepochka_obyazannostey/pkg"

func main() {
	device := &pkg.Device{Name: "Device-1"}
	updateSvc := &pkg.UpdateDataService{Name: "Update-1"}
	dataSvc := &pkg.DataService{Name: "-1"}

	device.SetNext(updateSvc)
	updateSvc.SetNext(dataSvc)

	data := &pkg.Data{}
	device.Execute(data)

}

//Тем самым мы реализовали паттерн "Цепочка обязанностей". Каждое звено решает какую-то свою задачу и передает после решения
//своей задачи следующему звену по цепочке.
