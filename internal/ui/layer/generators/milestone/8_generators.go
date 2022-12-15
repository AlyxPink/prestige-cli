package milestone

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type eight_generators struct {
	Milestone *layer.ModelMilestone
}

func Fetch8Generators(layers *layer.Layers) (milestone layer.Milestone) {
	model := eight_generators{
		Milestone: &layer.ModelMilestone{
			Name:        "8 Generators",
			Description: "Keep Prestige Upgrades on reset.",
			Layers:      layers,
		},
	}
	return &model
}

func (m *eight_generators) Tick() {
	if !m.Milestone.Reached && m.Done() {
		m.Milestone.Reached = true
	}
}

func (m *eight_generators) Done() bool {
	return m.Milestone.Layers.Generators.Model().AmountBest >= 8
}

func (m *eight_generators) Unlocked() bool {
	return true
}

func (m *eight_generators) Model() *layer.ModelMilestone {
	return m.Milestone
}
