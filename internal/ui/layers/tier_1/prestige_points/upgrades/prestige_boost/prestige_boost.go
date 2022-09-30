package prestige_boost

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
)

type PrestigeBoost struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func NewModel(pp *layers.Model, points *points.Points) PrestigeBoost {
	b := PrestigeBoost{
		Points:         points,
		PrestigePoints: pp,
		Upgrade: &upgrades.Model{
			Name:        "Prestige Boost",
			Description: "Prestige Points boost Point generation.",
			Unlocked:    true,
			Cost:        1,
		},
	}
	return b
}

func (pb *PrestigeBoost) Buy() {
	if pb.Upgrade.Enabled == true {
		return
	}
	if pb.PrestigePoints.Amount >= pb.Upgrade.Cost {
		pb.Upgrade.Enabled = true
		pb.PrestigePoints.Amount = pb.PrestigePoints.Amount - pb.Upgrade.Cost
	}
}

func (pb *PrestigeBoost) Tick() {
	pb.Points.Amount = pb.Points.Amount + (math.Pow(pb.PrestigePoints.Amount+2, 0.5) / 100)
}

func (pb *PrestigeBoost) GetModel() *upgrades.Model {
	return pb.Upgrade
}

func Fetch(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	upgradeModel := NewModel(layer, points)
	return &upgradeModel
}
