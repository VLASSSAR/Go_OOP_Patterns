package pkg

import "fmt"

type Device struct {
	Name string
	Next Service
}

func (device *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("The data from [%s] is received", device.Name)
		device.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device [%s].\n", device.Name)
	data.GetSource = true
	device.Next.Execute(data)
}

func (device *Device) SetNext(svc Service) {
	device.Next = svc
}
