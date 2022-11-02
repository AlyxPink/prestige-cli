package upgrade

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/upgrades"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
)

type prestigeBoost struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func FetchPrestigeBoost(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	model := prestigeBoost{
		Points:         points,
		PrestigePoints: layer,
		Upgrade: &upgrades.Model{
			Name:        "Prestige Boost",
			Description: "Prestige Points boost Point generation.",
			Cost:        1,
		},
	}
	return &model
}

func (model *prestigeBoost) Buy() {
	model.PrestigePoints.Amount = model.Upgrade.Buy(model.PrestigePoints.Amount)
}

func (model *prestigeBoost) Tick() {
	model.Points.Amount = model.Points.Amount + model.TickAmount()
}

func (model *prestigeBoost) Unlocked() bool {
	return model.PrestigePoints.Upgrades[0].GetModel().Enabled
}

func (model *prestigeBoost) TickAmount() float64 {
	return math.Pow(model.PrestigePoints.Amount+2, 0.5) / 100
}

func (model *prestigeBoost) GetModel() *upgrades.Model {
	return model.Upgrade
}
