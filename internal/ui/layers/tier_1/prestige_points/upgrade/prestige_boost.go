package upgrade

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
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
			Unlocked:    true,
			Cost:        1,
		},
	}
	return &model
}

func (pb *prestigeBoost) Buy() {
	pb.PrestigePoints.Amount = pb.Upgrade.Buy(pb.PrestigePoints.Amount)
}

func (pb *prestigeBoost) Tick() {
	pb.Points.Amount = pb.Points.Amount + pb.TickAmount()
}

func (pb *prestigeBoost) TickAmount() float64 {
	return math.Pow(pb.PrestigePoints.Amount+2, 0.5) / 100
}

func (pb *prestigeBoost) GetModel() *upgrades.Model {
	return pb.Upgrade
}
