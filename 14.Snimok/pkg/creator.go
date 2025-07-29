package pkg

type Creator struct {
	State string
}

func (c *Creator) Create() *Snapshot {
	return &Snapshot{state: c.State}
}

func (c *Creator) Restore(m *Snapshot) {
	c.State = m.GetSavedState()
}

func (c *Creator) SetState(state string) {
	c.State = state
}

func (c *Creator) GetState() string {
	return c.State
}
