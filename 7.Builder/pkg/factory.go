package pkg

type Factory struct {
	Collector Collector
}

// Так как наш завод может менять бренды в зависимости от ситуации, то мы должнысделать передачу интерфейса как аргумента, чтобы мы могли задавать
// любую комплектацию для нашего компьютера, который производится на этом заводе.
func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetMemory()
	factory.Collector.SetBrand()
	factory.Collector.SetGraphicCard()
	factory.Collector.SetMonitor()
	return factory.Collector.GetComputer()
}
