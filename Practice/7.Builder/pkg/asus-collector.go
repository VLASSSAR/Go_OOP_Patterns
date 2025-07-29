package pkg

type AsusCollector struct {
	Core         int
	Memory       int
	GraphicsCard int
	Monitor      int
	Brand        string
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetGraphicsCard() {
	collector.GraphicsCard = 1
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:         collector.Core,
		Memory:       collector.Memory,
		GraphicsCard: collector.GraphicsCard,
		Monitor:      collector.Monitor,
		Brand:        collector.Brand,
	}
}
