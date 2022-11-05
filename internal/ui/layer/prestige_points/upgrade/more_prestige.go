package upgrade

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer/upgrades"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
)

type morePrestige struct {
	Points         *points.Model
	PrestigePoints *layer.Model
	Upgrade        *upgrades.Model
}

func FetchMorePrestige(layer *layer.Model, points *points.Model) (upgrade upgrades.Upgrade) {
	model := morePrestige{
		Points:         points,
		PrestigePoints: layer,
		Upgrade: &upgrades.Model{
			Name:        "More Prestige",
			Description: "Prestige Point gain is increased by 80%.",
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

func (model *morePrestige) Effect() string {
	return ""
}

func (model *morePrestige) Unlocked() bool {
	return model.PrestigePoints.Upgrades[2].GetModel().Enabled
}

func (model *morePrestige) TickAmount() float64 {
	return 0
}

func (model *morePrestige) GetModel() *upgrades.Model {
	return model.Upgrade
}
