package upgrade

import (
	"fmt"
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type outerSynergy struct {
	Upgrade *layer.ModelUpgrade
}

func FetchOuterSynergy(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := outerSynergy{
		Upgrade: &layer.ModelUpgrade{
			Name:        "Outer Synergy",
			Description: "Self-Synergy is stronger based on your Generators.",
			Layers:      layers,
			Cost:        15,
		},
	}
	return &model
}

func (m *outerSynergy) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.Generators.Model())
}

func (m *outerSynergy) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *outerSynergy) Effect() string {
	return fmt.Sprintf("^%.2f", m.TickAmount())
}

func (m *outerSynergy) Unlocked() bool {
	return m.Upgrade.Layers.Generators.Model().Upgrades[2].Unlocked()
}

func (m *outerSynergy) TickAmount() float64 {
	var amount float64
	amount = m.Upgrade.Layers.Generators.Model().Amount
	amount = math.Sqrt(amount)
	amount = amount + 1
	if amount >= 400 {
		amount = math.Cbrt(amount)
		amount = amount * math.Pow(400, 2/3)
	}
	return amount
}

func (m *outerSynergy) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
