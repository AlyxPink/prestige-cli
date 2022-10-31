package upgrade

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
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
			Unlocked:    true,
			Cost:        5_000,
		},
	}
	return &model
}

func (rpb *reversePrestigeBoost) Buy() {
	rpb.PrestigePoints.Amount = rpb.Upgrade.Buy(rpb.PrestigePoints.Amount)
}

func (rpb *reversePrestigeBoost) Tick() {
	rpb.Points.Amount = rpb.Points.Amount + rpb.TickAmount()
}

func (rpb *reversePrestigeBoost) TickAmount() float64 {
	var amount float64
	amount = rpb.Points.Amount + 1
	amount = math.Log10(amount)
	amount = math.Cbrt(amount)
	amount = amount + 1
	return amount
}

func (rpb *reversePrestigeBoost) GetModel() *upgrades.Model {
	return rpb.Upgrade
}
