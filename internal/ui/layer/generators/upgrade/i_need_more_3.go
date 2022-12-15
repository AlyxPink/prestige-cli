package upgrade

import (
	"fmt"

	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type iNeedMore3 struct {
	Upgrade *layer.ModelUpgrade
}

func FetchINeedMore3(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := iNeedMore3{
		Upgrade: &layer.ModelUpgrade{
			Name:        "I Need More III",
			Description: "Generator Power boost its own generation.",
			Layers:      layers,
			Cost:        1e10,
		},
	}
	return &model
}

func (m *iNeedMore3) Buy() {
	m.Upgrade.Buy(m.Upgrade.Layers.Generators.Model())
}

func (m *iNeedMore3) Tick() {
	m.Upgrade.Layers.Points.Model().Amount = m.Upgrade.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *iNeedMore3) Effect() string {
	return fmt.Sprintf("%.2fx", m.TickAmount())
}

func (m *iNeedMore3) Unlocked() bool {
	return m.Upgrade.Layers.Generators.Model().Upgrades[4].Model().Enabled
}

func (m *iNeedMore3) TickAmount() float64 {
	var amount float64
	// TODO: power not implemented yet
	return amount
}

func (m *iNeedMore3) Model() *layer.ModelUpgrade {
	return m.Upgrade
}
