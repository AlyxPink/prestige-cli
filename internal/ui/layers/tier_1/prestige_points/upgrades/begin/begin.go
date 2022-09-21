package begin

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
)

type Begin struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func NewModel(pp *layers.Model, points *points.Points) Begin {
	b := Begin{
		Points:         points,
		PrestigePoints: pp,
		Upgrade: &upgrades.Model{
			Name:        "Begin",
			Description: "Generate 1 Point every second.",
			Unlocked:    true,
			Cost:        1,
		},
	}
	return b
}

func (b *Begin) Buy() {
	if b.PrestigePoints.Amount >= b.Upgrade.Cost {
		b.Upgrade.Enabled = true
		b.PrestigePoints.Amount = b.PrestigePoints.Amount - b.Upgrade.Cost
	}
}

func (b *Begin) Tick() {
	if !b.Upgrade.Enabled {
		return
	}
	b.Points.Amount = b.Points.Amount + 0.01
}

func Fetch(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	upgradeModel := NewModel(layer, points)
	return &upgradeModel
}
