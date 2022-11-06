package milestone

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type eight_boosters struct {
	Milestone *layer.ModelMilestone
}

func Fetch8Boosters(layers *layer.Layers) (milestone layer.Milestone) {
	model := eight_boosters{
		Milestone: &layer.ModelMilestone{
			Name:        "8 Boosters",
			Description: "Keep Prestige Upgrades on reset.",
			Layers:      layers,
		},
	}
	return &model
}

func (m *eight_boosters) Tick() {
	if !m.Milestone.Reached && m.Done() {
		m.Milestone.Reached = true
	}
}

func (m *eight_boosters) Done() bool {
	return m.Milestone.Layers.Boosters.Model().Amount >= 8
}

func (m *eight_boosters) Unlocked() bool {
	return true
}

func (m *eight_boosters) Model() *layer.ModelMilestone {
	return m.Milestone
}
