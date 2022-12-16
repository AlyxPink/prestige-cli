package milestone

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type fifteen_generators struct {
	Milestone *layer.ModelMilestone
}

func Fetch15Generators(layers *layer.Layers) (milestone layer.Milestone) {
	model := fifteen_generators{
		Milestone: &layer.ModelMilestone{
			Name:        "15 Generators",
			Description: "You can buy max Generators.",
			Layers:      layers,
		},
	}
	return &model
}

func (m *fifteen_generators) Tick() {
	if !m.Milestone.Reached && m.Done() {
		m.Milestone.Reached = true
	}
}

func (m *fifteen_generators) Done() bool {
	return m.Milestone.Layers.Generators.Model().AmountBest >= 15
}

func (m *fifteen_generators) Unlocked() bool {
	return true
}

func (m *fifteen_generators) Model() *layer.ModelMilestone {
	return m.Milestone
}
