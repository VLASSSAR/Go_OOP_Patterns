package pkg

type Factory struct {
	Collector Collector
}

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
	factory.Collector.SetMonitor()
	factory.Collector.SetGraphicsCard()
	return factory.Collector.GetComputer()

}
