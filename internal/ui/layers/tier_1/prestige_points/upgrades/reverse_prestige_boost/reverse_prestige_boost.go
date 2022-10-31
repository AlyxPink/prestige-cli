package reverse_prestige_boost

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
)

type ReversePrestigeBoost struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func NewModel(pp *layers.Model, points *points.Points) ReversePrestigeBoost {
	b := ReversePrestigeBoost{
		Points:         points,
		PrestigePoints: pp,
		Upgrade: &upgrades.Model{
			Name:        "Reverse Prestige Boost",
			Description: "Prestige Point gain is boosted by your Points.",
			Unlocked:    true,
			Cost:        5_000,
		},
	}
	return b
}

func (rpb *ReversePrestigeBoost) Buy() {
	rpb.PrestigePoints.Amount = rpb.Upgrade.Buy(rpb.PrestigePoints.Amount)
}

func (rpb *ReversePrestigeBoost) Tick() {
	rpb.Points.Amount = rpb.Points.Amount + rpb.TickAmount()
}

func (rpb *ReversePrestigeBoost) TickAmount() float64 {
	var amount float64
	amount = rpb.Points.Amount + 1
	amount = math.Log10(amount)
	amount = math.Cbrt(amount)
	amount = amount + 1
	return amount
}

func (rpb *ReversePrestigeBoost) GetModel() *upgrades.Model {
	return rpb.Upgrade
}

func Fetch(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	upgradeModel := NewModel(layer, points)
	return &upgradeModel
}
