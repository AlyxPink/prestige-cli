package upgrade

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type morePrestige struct {
	Upgrade *layer.ModelUpgrade
}

func FetchMorePrestige(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := morePrestige{
		Upgrade: &layer.ModelUpgrade{
			Name:        "More Prestige",
			Description: "Prestige Point gain is increased by 80%.",
			Layers:      layers,
			Cost:        20,
		},
	}
	return &model
}

func (m *morePrestige) Buy() {
	m.Upgrade.Layers.PrestigePoints.Model().Amount = m.Upgrade.Buy(m.Upgrade.Layers.PrestigePoints.Model().Amount)
}

func (m *morePrestige) Tick() {
}

func (m *morePrestige) Effect() string {
	return ""
}

func (m *morePrestige) Unlocked() bool {
	return m.Upgrade.Layers.PrestigePoints.Model().Upgrades[2].Model().Enabled
}

func (m *morePrestige) TickAmount() float64 {
	return 0
}

func (m *morePrestige) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
