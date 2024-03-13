package upgrade

import (
	"fmt"
	"math"

	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
)

type iNeedMore2 struct {
	Upgrade *layer.ModelUpgrade
}

func FetchINeedMore2(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := iNeedMore2{
		Upgrade: &layer.ModelUpgrade{
			Name:        "I Need More II",
			Description: "Best Prestige Points add to the Generator base.",
			Layers:      layers,
			Cost:        8,
		},
	}
	return &model
}

func (m *iNeedMore2) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.Generators.Model())
}

func (m *iNeedMore2) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *iNeedMore2) Effect() string {
	return fmt.Sprintf("+%.2f", m.TickAmount())
}

func (m *iNeedMore2) Unlocked() bool {
	return m.Upgrade.Layers.Generators.Model().AmountBest >= 8
}

func (m *iNeedMore2) TickAmount() float64 {
	var amount float64
	amount = m.Upgrade.Layers.PrestigePoints.Model().Amount
	amount = amount + 1
	amount = math.Log10(amount)
	amount = amount + 1
	amount = math.Log10(amount)
	amount = amount / 3
	return amount
}

func (m *iNeedMore2) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
