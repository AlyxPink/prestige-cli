package points

type Points struct {
	Amount float64
	countable
}

type countable interface {
	Count() float64
}

func (p *Points) Count() float64 {
	return p.Amount
}
