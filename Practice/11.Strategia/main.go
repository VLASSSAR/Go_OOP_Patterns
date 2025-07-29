package main

import "11.Strategia/pkg"

var (
	start      = 10
	end        = 100
	strategies = []pkg.Strategy{
		&pkg.PublicTransportStrategy{},
		&pkg.WalkStrategy{},
		&pkg.RoadStrategy{},
	}
)

func main() {
	nav := pkg.Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}
