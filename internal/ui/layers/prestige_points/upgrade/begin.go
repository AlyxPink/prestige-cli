package upgrade

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/upgrades"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
)

type begin struct {
	Points         *points.Model
	PrestigePoints *layers.Model
	Upgrade        *upgrades.Model
}

func FetchBegin(layer *layers.Model, points *points.Model) (upgrade upgrades.Upgrade) {
	model := begin{
		Points:         points,
		PrestigePoints: layer,
		Upgrade: &upgrades.Model{
			Name:        "Begin",
			Description: "Generate 1 Point every second.",
			Cost:        1,
		},
	}
	return &model
}

func (model *begin) Buy() {
	model.PrestigePoints.Amount = model.Upgrade.Buy(model.PrestigePoints.Amount)
}

func (model *begin) Tick() {
	model.Points.Amount = model.Points.Amount + model.TickAmount()/100
}

func (model *begin) Effect() string {
	return ""
}

func (model *begin) Unlocked() bool {
	return true
}

func (model *begin) TickAmount() float64 {
	return 1
}

func (model *begin) GetModel() *upgrades.Model {
	return model.Upgrade
}
