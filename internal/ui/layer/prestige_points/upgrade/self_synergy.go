package upgrade

import (
	"fmt"
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type selfSynergy struct {
	Upgrade *layer.ModelUpgrade
}

func FetchSelfSynergy(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := selfSynergy{
		Upgrade: &layer.ModelUpgrade{
			Name:        "Self-Synergy",
			Description: "Points boost their own generation.",
			Layers:      layers,
			Cost:        5,
		},
	}
	return &model
}

func (m *selfSynergy) Buy() {
	m.Upgrade.Layers.PrestigePoints.Model().Amount = m.Upgrade.Buy(m.Upgrade.Layers.PrestigePoints.Model().Amount)
}

func (m *selfSynergy) Tick() {
	m.Upgrade.Layers.Points.Amount = m.Upgrade.Layers.Points.Amount + m.TickAmount()/100
}

func (m *selfSynergy) Effect() string {
	return fmt.Sprintf("%.2fx", m.TickAmount())
}

func (m *selfSynergy) Unlocked() bool {
	return m.Upgrade.Layers.PrestigePoints.Model().Upgrades[1].GetModel().Enabled
}

func (m *selfSynergy) TickAmount() float64 {
	var amount float64
	amount = m.Upgrade.Layers.Points.Amount + 1
	amount = math.Log10(amount)
	amount = math.Pow(amount, 0.75)
	amount = amount + 1
	return amount
}

func (m *selfSynergy) GetModel() *layer.ModelUpgrade {
	return m.Upgrade
}
