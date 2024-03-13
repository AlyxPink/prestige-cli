package upgrade

import (
	"fmt"
	"math"

	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
)

type iNeedMore1 struct {
	Upgrade *layer.ModelUpgrade
}

func FetchINeedMore1(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := iNeedMore1{
		Upgrade: &layer.ModelUpgrade{
			Name:        "I Need More!",
			Description: "Boosters add to the Generator base.",
			Layers:      layers,
			Cost:        7,
		},
	}
	return &model
}

func (m *iNeedMore1) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.Generators.Model())
}

func (m *iNeedMore1) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *iNeedMore1) Effect() string {
	return fmt.Sprintf("+%.2f", m.TickAmount())
}

func (m *iNeedMore1) Unlocked() bool {
	return true
}

func (m *iNeedMore1) TickAmount() float64 {
	var amount float64
	amount = m.Upgrade.Layers.Boosters.Model().Amount
	amount = amount + 1
	amount = math.Log10(amount)
	amount = math.Sqrt(amount)
	amount = amount / 3
	return amount
}

func (m *iNeedMore1) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
