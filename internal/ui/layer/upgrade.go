package layer

type ModelUpgrade struct {
	Name        string
	Description string
	Amount      uint
	Cost        float64
	Enabled     bool
	Layers      *Layers
}

type Upgrade interface {
	Tick()
	TickAmount() float64
	Effect() string

	Buy()

	Unlocked() bool

	Model() *ModelUpgrade
}

func (m *ModelUpgrade) Buy(currency *Model) {
	if !m.Enabled && currency.Amount >= m.Cost {
		m.Enabled = true
		currency.Amount = currency.Amount - m.Cost
	}
}
