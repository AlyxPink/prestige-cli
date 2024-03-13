package achievement

import (
	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
)

type allProgressGone struct {
	Achievement *layer.ModelAchievement
}

func FetchAllProgressGone(layers *layer.Layers) (Achievement layer.Achievement) {
	model := allProgressGone{
		Achievement: &layer.ModelAchievement{
			Name:        "All that progress is gone!",
			Description: "Perform a Prestige reset.",
			Layers:      layers,
		},
	}
	return &model
}

func (m *allProgressGone) Tick() {
	if !m.Achievement.Achieved && m.Done() {
		m.Achievement.Achieved = true
	}
}

func (m *allProgressGone) Done() bool {
	return m.Model().Layers.PrestigePoints.Model().Amount > 0
}

func (m *allProgressGone) Effect() {}

func (m *allProgressGone) Unlocked() bool {
	return true
}

func (m *allProgressGone) Model() *layer.ModelAchievement {
	return m.Achievement
}
