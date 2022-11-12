package upgrade

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type bpCombo struct {
	Upgrade *layer.ModelUpgrade
}

func FetchBPCombo(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := bpCombo{
		Upgrade: &layer.ModelUpgrade{
			Name:        "BP Combo",
			Description: "Best Boosters boost Prestige Point gain.",
			Layers:      layers,
			Cost:        3,
		},
	}
	return &model
}

func (m *bpCombo) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.Boosters.Model())
}

func (m *bpCombo) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *bpCombo) Effect() string {
	return ""
}

func (m *bpCombo) Unlocked() bool {
	return true
}

func (m *bpCombo) TickAmount() float64 {
	var amount float64
	amount = math.Sqrt(m.Upgrade.Layers.Boosters.Model().AmountBest)
	amount = amount + 1
	return amount
}

func (m *bpCombo) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
