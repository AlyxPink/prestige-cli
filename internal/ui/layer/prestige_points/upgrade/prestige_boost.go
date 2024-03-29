package upgrade

import (
	"fmt"
	"math"

	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
)

type prestigeBoost struct {
	Upgrade *layer.ModelUpgrade
}

func FetchPrestigeBoost(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := prestigeBoost{
		Upgrade: &layer.ModelUpgrade{
			Name:        "Prestige Boost",
			Description: "Prestige Points boost Point generation.",
			Layers:      layers,
			Cost:        1,
		},
	}
	return &model
}

func (m *prestigeBoost) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.PrestigePoints.Model())
}

func (m *prestigeBoost) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *prestigeBoost) Effect() string {
	return fmt.Sprintf("%.2fx", m.TickAmount())
}

func (m *prestigeBoost) Unlocked() bool {
	return m.Upgrade.Layers.PrestigePoints.Model().Upgrades[0].Model().Enabled
}

func (m *prestigeBoost) TickAmount() float64 {
	var amount float64
	amount = m.Upgrade.Layers.PrestigePoints.Model().Amount
	amount = amount + 2
	amount = math.Pow(amount, 0.5)
	if m.Upgrade.Layers.Generators.Model().Upgrades[3].Model().Enabled { // Boost the Boost
		amount = math.Pow(amount, 1.5)
	}
	return amount
}

func (m *prestigeBoost) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
