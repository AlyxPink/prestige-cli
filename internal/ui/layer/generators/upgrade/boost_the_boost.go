package upgrade

import (
	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
)

type boostTheBoost struct {
	Upgrade *layer.ModelUpgrade
}

func FetchBoostTheBoost(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := boostTheBoost{
		Upgrade: &layer.ModelUpgrade{
			Name:        "Boost the Boost",
			Description: "Prestige Boost is raised to the power of 1.5.",
			Layers:      layers,
			Cost:        13,
		},
	}
	return &model
}

func (m *boostTheBoost) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.Generators.Model())
}

func (m *boostTheBoost) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *boostTheBoost) Effect() string {
	return ""
}

func (m *boostTheBoost) Unlocked() bool {
	return m.Upgrade.Layers.Generators.Model().AmountBest >= 10
}

func (m *boostTheBoost) TickAmount() float64 {
	var amount float64
	return amount
}

func (m *boostTheBoost) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
