package upgrade

import (
	"fmt"
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
			Cost:        5,
		},
	}
	return &model
}

func (model *selfSynergy) Buy() {
	model.PrestigePoints.Amount = model.Upgrade.Buy(model.PrestigePoints.Amount)
}

func (model *selfSynergy) Tick() {
	model.Points.Amount = model.Points.Amount + model.TickAmount()
}

func (model *selfSynergy) Effect() string {
	return fmt.Sprintf("%.2fx", model.TickAmount()*100)
}

func (model *selfSynergy) Unlocked() bool {
	return model.PrestigePoints.Upgrades[1].GetModel().Enabled
}

func (model *selfSynergy) TickAmount() float64 {
	var amount float64
	amount = model.Points.Amount + 1
	amount = math.Log10(amount)
	amount = math.Pow(amount, 0.75)
	amount = amount + 1
	amount = amount / 100
	return amount
}

func (model *selfSynergy) GetModel() *upgrades.Model {
	return model.Upgrade
}
