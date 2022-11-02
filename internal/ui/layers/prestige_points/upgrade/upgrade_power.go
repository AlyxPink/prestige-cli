package upgrade

import (
	"fmt"
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
			Cost:        75,
		},
	}
	return &model
}

func (model *upgradePower) Buy() {
	model.PrestigePoints.Amount = model.Upgrade.Buy(model.PrestigePoints.Amount)
}

func (model *upgradePower) Tick() {
	model.Points.Amount = model.Points.Amount + model.TickAmount()
}

func (model *upgradePower) Effect() string {
	return fmt.Sprintf("%.2fx", model.TickAmount())
}

func (model *upgradePower) Unlocked() bool {
	return model.PrestigePoints.Upgrades[2].GetModel().Enabled
}

func (model *upgradePower) TickAmount() float64 {
	var amount float64
	upgrades_enabled_count := len(model.PrestigePoints.ListUpgradeEnabled())
	amount = math.Pow(1.4, float64(upgrades_enabled_count))
	return amount
}

func (model *upgradePower) GetModel() *upgrades.Model {
	return model.Upgrade
}
