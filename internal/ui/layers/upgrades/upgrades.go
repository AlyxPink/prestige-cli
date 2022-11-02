package upgrades

type Model struct {
	Name        string
	Description string
	Amount      uint
	Cost        float64
	Enabled     bool
	Unlocked    bool
}

type Upgrade interface {
	Tick()
	TickAmount() float64
	Buy()
	GetModel() *Model
}

func (m *Model) Buy(currency float64) float64 {
	if !m.Enabled && currency >= m.Cost {
		m.Enabled = true
		return currency - m.Cost
	}
	return currency
}
