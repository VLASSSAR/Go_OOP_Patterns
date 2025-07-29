package pkg

type Snapshot struct {
	state string
}

func (m *Snapshot) GetSavedState() string {
	return m.state
}
