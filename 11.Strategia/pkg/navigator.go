package pkg

// Создадим сущность навигатор, к которому будут применяться стратегии.
type Navigator struct {
	Strategy
}

func (nav *Navigator) SetStrategy(str Strategy) {
	nav.Strategy = str
}
