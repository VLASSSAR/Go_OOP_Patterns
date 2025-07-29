package pkg

type HpCollector struct {
	Core         int
	Memory       int
	GraphicsCard int
	Monitor      int
	Brand        string
}

func (collector *HpCollector) SetCore() {
	collector.Core = 30
}

func (collector *HpCollector) SetGraphicsCard() {
	collector.GraphicsCard = 2
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 2
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 256
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:         collector.Core,
		Memory:       collector.Memory,
		GraphicsCard: collector.GraphicsCard,
		Monitor:      collector.Monitor,
		Brand:        collector.Brand,
	}
}
