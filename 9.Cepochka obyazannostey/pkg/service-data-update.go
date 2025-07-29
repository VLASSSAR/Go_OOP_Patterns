package pkg

import "fmt"

type UpdateDevice struct {
	Name string
	Next Service
}

func (upd *UpdateDevice) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("The data from [%s] is updated\n", upd.Name)
		upd.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from device [%s].\n", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDevice) SetNext(svc Service) {
	upd.Next = svc
}
