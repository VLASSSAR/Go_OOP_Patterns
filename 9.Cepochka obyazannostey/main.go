package main

import "9.Cepochka_obyazannostey/pkg"

func main() {
	device := &pkg.Device{Name: "DeviceService-1"}
	updateSvc := &pkg.UpdateDevice{Name: "UpdateService-1"}
	dataSvc := &pkg.DataService{Name: "DataService-1"}

	device.SetNext(updateSvc)
	updateSvc.SetNext(dataSvc)

	data := &pkg.Data{}
	device.Execute(data)

}
