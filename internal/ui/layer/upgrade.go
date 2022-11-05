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

func (m *ModelUpgrade) Buy(currency float64) float64 {
	if !m.Enabled && currency >= m.Cost {
		m.Enabled = true
		return currency - m.Cost
	}
	return currency
}
