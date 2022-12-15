package upgrade

import (
	"fmt"
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type gpCombo struct {
	Upgrade *layer.ModelUpgrade
}

func FetchGPCombo(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := gpCombo{
		Upgrade: &layer.ModelUpgrade{
			Name:        "GP Combo",
			Description: "Best Generators boost Prestige Point gain.",
			Layers:      layers,
			Cost:        3,
		},
	}
	return &model
}

func (m *gpCombo) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.Generators.Model())
}

func (m *gpCombo) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *gpCombo) Effect() string {
	return fmt.Sprintf("%.2fx", m.TickAmount())
}

func (m *gpCombo) Unlocked() bool {
	return true
}

func (m *gpCombo) TickAmount() float64 {
	var amount float64
	amount = math.Sqrt(m.Upgrade.Layers.Generators.Model().AmountBest)
	amount = amount + 1
	return amount
}

func (m *gpCombo) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
