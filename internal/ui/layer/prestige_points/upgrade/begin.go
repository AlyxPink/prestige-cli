package upgrade

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type begin struct {
	Upgrade *layer.ModelUpgrade
}

func FetchBegin(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := begin{
		Upgrade: &layer.ModelUpgrade{
			Name:        "Begin",
			Description: "Generate 1 Point every second.",
			Layers:      layers,
			Cost:        1,
		},
	}
	return &model
}

func (m *begin) Buy() {
	m.Upgrade.Layers.PrestigePoints.Model().Amount = m.Upgrade.Buy(m.Upgrade.Layers.PrestigePoints.Model().Amount)
}

func (m *begin) Tick() {
	m.Upgrade.Layers.Points.Amount = m.Upgrade.Layers.Points.Amount + m.TickAmount()/100
}

func (m *begin) Effect() string {
	return ""
}

func (m *begin) Unlocked() bool {
	return true
}

func (m *begin) TickAmount() float64 {
	return 1
}

func (m *begin) GetModel() *layer.ModelUpgrade {
	return m.Upgrade
}
