package upgrade

import (
	"fmt"
	"math"

	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
)

type reversePrestigeBoost struct {
	Upgrade *layer.ModelUpgrade
}

func FetchReversePrestigeBoost(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := reversePrestigeBoost{
		Upgrade: &layer.ModelUpgrade{
			Name:        "Reverse Prestige Boost",
			Description: "Prestige Point gain is boosted by your Points.",
			Layers:      layers,
			Cost:        5_000,
		},
	}
	return &model
}

func (m *reversePrestigeBoost) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.PrestigePoints.Model())
}

func (m *reversePrestigeBoost) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *reversePrestigeBoost) Effect() string {
	return fmt.Sprintf("%.2fx", m.TickAmount())
}

func (m *reversePrestigeBoost) Unlocked() bool {
	return m.Upgrade.Layers.PrestigePoints.Model().Upgrades[2].Model().Enabled
}

func (m *reversePrestigeBoost) TickAmount() float64 {
	var amount float64
	amount = m.Upgrade.Layers.Points.Model().Amount + 1
	amount = math.Log10(amount)
	amount = math.Cbrt(amount)
	amount = amount + 1
	return amount
}

func (m *reversePrestigeBoost) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
