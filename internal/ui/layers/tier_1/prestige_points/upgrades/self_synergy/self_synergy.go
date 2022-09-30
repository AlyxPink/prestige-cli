package self_synergy

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
)

type SelfSynergy struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func NewModel(pp *layers.Model, points *points.Points) SelfSynergy {
	b := SelfSynergy{
		Points:         points,
		PrestigePoints: pp,
		Upgrade: &upgrades.Model{
			Name:        "Self-Synergy",
			Description: "Points boost their own generation.",
			Unlocked:    true,
			Cost:        5,
		},
	}
	return b
}

func (ss *SelfSynergy) Buy() {
	if ss.Upgrade.Enabled == true {
		return
	}
	if ss.PrestigePoints.Amount >= ss.Upgrade.Cost {
		ss.Upgrade.Enabled = true
		ss.PrestigePoints.Amount = ss.PrestigePoints.Amount - ss.Upgrade.Cost
	}
}

func (ss *SelfSynergy) Tick() {
	ss.Points.Amount = ss.Points.Amount + math.Log10(math.Pow(ss.Points.Amount+1, 0.75))/100
}

func (ss *SelfSynergy) GetModel() *upgrades.Model {
	return ss.Upgrade
}

func Fetch(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	upgradeModel := NewModel(layer, points)
	return &upgradeModel
}
