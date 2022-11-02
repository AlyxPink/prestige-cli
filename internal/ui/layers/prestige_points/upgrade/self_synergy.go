package upgrade

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/upgrades"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
)

type selfSynergy struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func FetchSelfSynergy(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	model := selfSynergy{
		Points:         points,
		PrestigePoints: layer,
		Upgrade: &upgrades.Model{
			Name:        "Self-Synergy",
			Description: "Points boost their own generation.",
			Unlocked:    true,
			Cost:        5,
		},
	}
	return &model
}

func (ss *selfSynergy) Buy() {
	ss.PrestigePoints.Amount = ss.Upgrade.Buy(ss.PrestigePoints.Amount)
}

func (ss *selfSynergy) Tick() {
	ss.Points.Amount = ss.Points.Amount + ss.TickAmount()
}

func (ss *selfSynergy) TickAmount() float64 {
	return math.Log10(math.Pow(ss.Points.Amount+1, 0.75)) / 100
}

func (ss *selfSynergy) GetModel() *upgrades.Model {
	return ss.Upgrade
}
