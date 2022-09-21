package begin

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
)

type Begin struct {
	Points  *points.Points
	Upgrade *upgrades.Model
}

func NewModel(points *points.Points) Begin {
	b := Begin{
		Points: points,
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
	if b.Points.Amount >= b.Upgrade.Cost {
		b.Upgrade.Enabled = true
		b.Points.Amount = b.Points.Amount - b.Upgrade.Cost
	}
}

func (b *Begin) Tick() {
	if !b.Upgrade.Enabled {
		return
	}
	b.Points.Amount = b.Points.Amount + 0.01
}

func Fetch(points *points.Points) (layer upgrades.Upgrade) {
	upgradeModel := NewModel(points)
	return &upgradeModel
}
