package upgrade

import (
	"fmt"
	"math"

	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
)

type upgradePower struct {
	Upgrade *layer.ModelUpgrade
}

func FetchUpgradePower(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := upgradePower{
		Upgrade: &layer.ModelUpgrade{
			Name:        "Upgrade Power",
			Description: "Point generation is faster based on your Prestige Upgrades bought.",
			Layers:      layers,
			Cost:        75,
		},
	}
	return &model
}

func (m *upgradePower) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.PrestigePoints.Model())
}

func (m *upgradePower) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *upgradePower) Effect() string {
	return fmt.Sprintf("%.2fx", m.TickAmount())
}

func (m *upgradePower) Unlocked() bool {
	return m.Upgrade.Layers.PrestigePoints.Model().Upgrades[2].Model().Enabled
}

func (m *upgradePower) TickAmount() float64 {
	var amount float64
	upgrades_enabled_count := len(m.Upgrade.Layers.PrestigePoints.Model().ListUpgradeEnabled())
	amount = math.Pow(1.4, float64(upgrades_enabled_count))
	return amount
}

func (m *upgradePower) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
