package milestone

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type fifteen_boosters struct {
	Milestone *layer.ModelMilestone
}

func Fetch15Boosters(layers *layer.Layers) (milestone layer.Milestone) {
	model := fifteen_boosters{
		Milestone: &layer.ModelMilestone{
			Name:        "15 Boosters",
			Description: "You can buy max Boosters.",
			Layers:      layers,
		},
	}
	return &model
}

func (m *fifteen_boosters) Tick() {
	if !m.Milestone.Reached && m.Done() {
		m.Milestone.Reached = true
	}
}

func (m *fifteen_boosters) Done() bool {
	return m.Milestone.Layers.Boosters.Model().Amount >= 15
}

func (m *fifteen_boosters) Unlocked() bool {
	return true
}

func (m *fifteen_boosters) Model() *layer.ModelMilestone {
	return m.Milestone
}
