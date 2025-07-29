package pkg

import "fmt"

type DataService struct {
	Name string
	Next Service
}

func (dataSvc *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("The data is not updated")
		return
	}
	fmt.Println("Data save")
}

func (dataSvc *DataService) SetNext(svc Service) {
	dataSvc.Next = svc
}
