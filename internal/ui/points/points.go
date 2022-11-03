package points

type Model struct {
	Amount float64
	countable
}

type countable interface {
	Count() float64
}

func (m *Model) Count() float64 {
	return m.Amount
}
