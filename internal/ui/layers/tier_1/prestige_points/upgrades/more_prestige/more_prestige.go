package more_prestige

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
)

type MorePrestige struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func NewModel(pp *layers.Model, points *points.Points) MorePrestige {
	mp := MorePrestige{
		Points:         points,
		PrestigePoints: pp,
		Upgrade: &upgrades.Model{
			Name:        "More Prestige",
			Description: "Prestige Point gain is increased by 80%.",
			Unlocked:    true,
			Cost:        20,
		},
	}
	return mp
}

func (mp *MorePrestige) Buy() {
	mp.PrestigePoints.Amount = mp.Upgrade.Buy(mp.PrestigePoints.Amount)
}

func (mp *MorePrestige) Tick() {
}

func (mp *MorePrestige) TickAmount() float64 {
	return 0
}

func (mp *MorePrestige) GetModel() *upgrades.Model {
	return mp.Upgrade
}

func Fetch(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	upgradeModel := NewModel(layer, points)
	return &upgradeModel
}
