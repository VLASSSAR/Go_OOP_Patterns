package main

import "12.Nabludatel/pkg"

func main() {
	sub1 := &pkg.Subscriber{Name: "Sub-1"}
	sub2 := &pkg.Subscriber{Name: "Sub-2"}
	sub3 := &pkg.Subscriber{Name: "Sub-3"}
	subN := &pkg.Subscriber{Name: "Sub-N"}

	channel := pkg.Publisher{
		Name:      "Publisher top",
		Consumers: map[string]pkg.Consumer{},
	}
	channel.Subscribe(sub1)
	channel.Subscribe(sub2)
	channel.Subscribe(sub3)
	channel.Subscribe(subN)
	channel.Notify()

	channel.UnSubscribe(sub2)
	channel.Notify()

}
