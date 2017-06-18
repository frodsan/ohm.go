package rom

type Model struct {
	Id string
}

func (m Model) IsNew() bool {
	return len(m.Id) == 0
}

func (m Model) IsPersisted() bool {
	return !m.IsNew()
}
