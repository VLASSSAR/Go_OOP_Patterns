package pkg

type Guardian struct {
	Items []*Snapshot
}

func (g *Guardian) Add(m *Snapshot) {
	g.Items = append(g.Items, m)
}

func (g *Guardian) Get(index int) *Snapshot {
	return g.Items[index]
}
