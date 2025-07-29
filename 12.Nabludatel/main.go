package main

import "12.Nabludatel/pkg"

func main() {
	sub1 := &pkg.Subscriber{Name: "Sub"}
	sub2 := &pkg.Subscriber{Name: "Sub-2"}
	sub3 := &pkg.Subscriber{Name: "Sub-3"}
	subN := &pkg.Subscriber{Name: "Sub-N"}
	channel := pkg.Publisher{
		Name:      "Publisher channel",
		Consumers: map[string]pkg.Consumer{},
	}
	channel.Subscribe(sub1)
	channel.Subscribe(sub2)
	channel.Subscribe(sub3)
	channel.Subscribe(subN)
	channel.Notify()
	println("Unsubscribe Sub-3")
	channel.Unsubscribe(sub3)
	channel.Notify()
}
