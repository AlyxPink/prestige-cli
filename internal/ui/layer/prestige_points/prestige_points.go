package prestige_points

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer/prestige_points/achievement"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer/prestige_points/upgrade"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	layer *layer.Model
}

func NewModel(id int, layers *layer.Layers, ctx *context.ProgramContext) Model {
	m := Model{
		layer: &layer.Model{
			Name:     "Prestige Points",
			Id:       id,
			Tier:     1,
			Unlocked: true,
			Layers:   layers,
			Ctx:      ctx,
		},
	}

	m.layer.Upgrades = []layer.Upgrade{
		upgrade.FetchBegin(layers),
		upgrade.FetchPrestigeBoost(layers),
		upgrade.FetchSelfSynergy(layers),
		upgrade.FetchMorePrestige(layers),
		upgrade.FetchUpgradePower(layers),
		upgrade.FetchReversePrestigeBoost(layers),
	}

	m.layer.Achievements = []layer.Achievement{
		achievement.FetchAllProgressGone(layers),
		achievement.FetchPointHog(layers),
		achievement.FetchPrestigeAllTheWay(layers),
		achievement.FetchPrestigeExp2(layers),
	}

	return m
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.layer.UpdateProgramContext(ctx)
}

func (m *Model) Model() *layer.Model {
	return m.layer
}

func (m *Model) Unlocked() bool {
	return true
}

func (m *Model) Tick() {}

func (m *Model) TickAmount() float64 {
	// Return 0.0 if nothing is being generated
	if !m.layer.Upgrades[0].Model().Enabled {
		return 0.0
	}
	amount := 1.0
	for _, upgrade := range m.layer.Upgrades {
		if upgrade.Model().Enabled {
			amount = amount * upgrade.TickAmount()
		}
	}
	return amount
}

func (m *Model) Prestige() {
	if m.PrestigeAmount() < 1 {
		return
	}
	m.layer.Amount = m.layer.Amount + m.PrestigeAmount()
	m.layer.AmountTotal = m.layer.AmountTotal + m.PrestigeAmount()
	m.layer.SaveBestAmount()
	m.layer.Layers.Points.Model().Amount = 0
}

func (m *Model) PrestigeAmount() float64 {
	if m.layer.Layers.Points.Model().Amount < m.PrestigeRequirement() {
		return 0
	}
	gain := m.layer.Layers.Points.Model().Amount / m.PrestigeRequirement()
	gain = math.Pow(gain, 0.5)
	gain = gain * m.GainMult()
	gain = math.Pow(gain, m.GainExp())
	return gain
}

func (m *Model) PrestigeRequirement() float64 {
	return 10
}

func (m *Model) GainMult() float64 {
	mult := 1.0
	if m.layer.Upgrades[3].Model().Enabled { // If "more_prestige" upgrade enabled
		mult = mult * 1.8
	}
	if m.layer.Achievements[2].Model().Achieved { // If "Prestige all the Way" achievement achieved
		mult = mult * 1.1
	}
	return mult
}

func (m *Model) GainExp() float64 {
	return 1
}

func (m *Model) Reset() {
	m.layer.Amount = 0
	for _, upgrade := range m.layer.Upgrades {
		upgrade.Model().Enabled = false
	}
}

func (m Model) Update(msg tea.Msg) (layer.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func Fetch(id int, layers *layer.Layers, ctx context.ProgramContext) (layer layer.Layer) {
	layerModel := NewModel(id, layers, &ctx)
	return &layerModel
}
