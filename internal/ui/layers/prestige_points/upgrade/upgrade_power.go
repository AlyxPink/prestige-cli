package upgrade

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/upgrades"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
)

type upgradePower struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func FetchUpgradePower(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	model := upgradePower{
		Points:         points,
		PrestigePoints: layer,
		Upgrade: &upgrades.Model{
			Name:        "Upgrade Power",
			Description: "Point generation is faster based on your Prestige Upgrades bought.",
			Unlocked:    true,
			Cost:        75,
		},
	}
	return &model
}

func (up *upgradePower) Buy() {
	up.PrestigePoints.Amount = up.Upgrade.Buy(up.PrestigePoints.Amount)
}

func (up *upgradePower) Tick() {
	up.Points.Amount = up.Points.Amount + up.TickAmount()
}

func (up *upgradePower) TickAmount() float64 {
	var amount float64
	upgrades_enabled_count := len(up.PrestigePoints.ListUpgradeEnabled())
	amount = math.Pow(1.4, float64(upgrades_enabled_count))
	return amount
}

func (up *upgradePower) GetModel() *upgrades.Model {
	return up.Upgrade
}
