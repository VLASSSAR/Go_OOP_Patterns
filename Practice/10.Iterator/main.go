package main

import (
	"10.Iterator/pkg"
	"fmt"
)

var routes = pkg.Routes{
	Routes: []pkg.Route{
		{Name: "Маршрут-1", TravelTime: 100},
		{Name: "Маршрут-2", TravelTime: 14},
		{Name: "Маршрут-3", TravelTime: 654},
		{Name: "Маршрут-4", TravelTime: 124},
		{Name: "Маршрут-5", TravelTime: 111},
	},
}

func main() {
	for routes.HasNext() {
		route := routes.GetNext()
		fmt.Printf("R:[%s], Time:[%d]\n", route.Name, route.TravelTime)
	}
}
