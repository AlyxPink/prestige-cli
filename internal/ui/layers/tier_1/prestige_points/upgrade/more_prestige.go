package upgrade

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
)

type morePrestige struct {
	Points         *points.Points
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func FetchMorePrestige(layer *layers.Model, points *points.Points) (upgrade upgrades.Upgrade) {
	model := morePrestige{
		Points:         points,
		PrestigePoints: layer,
		Upgrade: &upgrades.Model{
			Name:        "More Prestige",
			Description: "Prestige Point gain is increased by 80%.",
			Unlocked:    true,
			Cost:        20,
		},
	}
	return &model
}

func (model *morePrestige) Buy() {
	model.PrestigePoints.Amount = model.Upgrade.Buy(model.PrestigePoints.Amount)
}

func (model *morePrestige) Tick() {
}

func (model *morePrestige) TickAmount() float64 {
	return 0
}

func (model *morePrestige) GetModel() *upgrades.Model {
	return model.Upgrade
}
