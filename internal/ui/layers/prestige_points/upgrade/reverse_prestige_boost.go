package upgrade

import (
	"fmt"
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/upgrades"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
)

type reversePrestigeBoost struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func FetchReversePrestigeBoost(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	model := reversePrestigeBoost{
		Points:         points,
		PrestigePoints: layer,
		Upgrade: &upgrades.Model{
			Name:        "Reverse Prestige Boost",
			Description: "Prestige Point gain is boosted by your Points.",
			Cost:        5_000,
		},
	}
	return &model
}

func (model *reversePrestigeBoost) Buy() {
	model.PrestigePoints.Amount = model.Upgrade.Buy(model.PrestigePoints.Amount)
}

func (model *reversePrestigeBoost) Tick() {
	model.Points.Amount = model.Points.Amount + model.TickAmount()
}

func (model *reversePrestigeBoost) Effect() string {
	return fmt.Sprintf("%.2fx", model.TickAmount())
}

func (model *reversePrestigeBoost) Unlocked() bool {
	return model.PrestigePoints.Upgrades[2].GetModel().Enabled
}

func (model *reversePrestigeBoost) TickAmount() float64 {
	var amount float64
	amount = model.Points.Amount + 1
	amount = math.Log10(amount)
	amount = math.Cbrt(amount)
	amount = amount + 1
	return amount
}

func (model *reversePrestigeBoost) GetModel() *upgrades.Model {
	return model.Upgrade
}
