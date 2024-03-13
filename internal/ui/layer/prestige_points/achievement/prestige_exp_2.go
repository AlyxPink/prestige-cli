package achievement

import (
	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
)

type prestigeExp2 struct {
	Achievement *layer.ModelAchievement
}

func FetchPrestigeExp2(layers *layer.Layers) (Achievement layer.Achievement) {
	model := prestigeExp2{
		Achievement: &layer.ModelAchievement{
			Name:        "Prestige^2",
			Description: "Reach 25 Prestige Points.",
			Layers:      layers,
		},
	}
	return &model
}

func (m *prestigeExp2) Tick() {
	if !m.Achievement.Achieved && m.Done() {
		m.Achievement.Achieved = true
	}
}

func (m *prestigeExp2) Done() bool {
	if m.Model().Layers.PrestigePoints.Model().Amount > 25 {
		return true
	}
	return m.Model().Layers.PrestigePoints.Model().Amount > 25
}

func (m *prestigeExp2) Effect() {}

func (m *prestigeExp2) Unlocked() bool {
	return true
}

func (m *prestigeExp2) Model() *layer.ModelAchievement {
	return m.Achievement
}
