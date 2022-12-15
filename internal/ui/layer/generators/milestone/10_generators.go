package milestone

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type ten_generators struct {
	Milestone *layer.ModelMilestone
}

func Fetch10Generators(layers *layer.Layers) (milestone layer.Milestone) {
	model := ten_generators{
		Milestone: &layer.ModelMilestone{
			Name:        "8 Generators",
			Description: "You gain 100% of Prestige Point gain every second.",
			Layers:      layers,
		},
	}
	return &model
}

func (m *ten_generators) Tick() {
	if !m.Milestone.Reached && m.Done() {
		m.Milestone.Reached = true
	}
}

func (m *ten_generators) Done() bool {
	return m.Milestone.Layers.Generators.Model().AmountBest >= 10
}

func (m *ten_generators) Unlocked() bool {
	return true
}

func (m *ten_generators) Model() *layer.ModelMilestone {
	return m.Milestone
}
