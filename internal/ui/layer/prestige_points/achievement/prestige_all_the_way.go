package achievement

import (
	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
)

type prestigeAllTheWay struct {
	Achievement *layer.ModelAchievement
}

func FetchPrestigeAllTheWay(layers *layer.Layers) (Achievement layer.Achievement) {
	model := prestigeAllTheWay{
		Achievement: &layer.ModelAchievement{
			Name:        "Prestige all the Way",
			Description: "Purchase 3 Prestige Upgrades. \nReward: Gain 10% more Prestige Points.",
			Layers:      layers,
		},
	}
	return &model
}

func (m *prestigeAllTheWay) Tick() {
	if !m.Achievement.Achieved && m.Done() {
		m.Achievement.Achieved = true
	}
}

func (m *prestigeAllTheWay) Done() bool {
	return len(m.Model().Layers.PrestigePoints.Model().ListUpgradeEnabled()) >= 3
}

func (m *prestigeAllTheWay) Effect() {}

func (m *prestigeAllTheWay) Unlocked() bool {
	return true
}

func (m *prestigeAllTheWay) Model() *layer.ModelAchievement {
	return m.Achievement
}
