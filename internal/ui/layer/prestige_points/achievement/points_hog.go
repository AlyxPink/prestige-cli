package achievement

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type pointHog struct {
	Achievement *layer.ModelAchievement
}

func FetchPointHog(layers *layer.Layers) (Achievement layer.Achievement) {
	model := pointHog{
		Achievement: &layer.ModelAchievement{
			Name:        "Point Hog",
			Description: "Reach 25 Points.",
			Layers:      layers,
		},
	}
	return &model
}

func (m *pointHog) Tick() {
	if !m.Achievement.Achieved && m.Done() {
		m.Achievement.Achieved = true
	}
}

func (m *pointHog) Done() bool {
	return m.Model().Layers.Points.Model().Amount >= 25
}

func (m *pointHog) Effect() {}

func (m *pointHog) Unlocked() bool {
	return true
}

func (m *pointHog) Model() *layer.ModelAchievement {
	return m.Achievement
}
