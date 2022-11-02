package generators

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/upgrades"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	layer    *layers.Model
	upgrades []upgrades.Upgrade
}

func NewModel(id int, ctx *context.ProgramContext) Model {
	g := Model{
		layer: &layers.Model{
			Id:   id,
			Tier: 2,
			Ctx:  ctx,
			Name: "Generators",
		},
	}

	return g
}

func (m *Model) Model() *layers.Model {
	return m.layer
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.layer.UpdateProgramContext(ctx)
}

func (m *Model) Tick() {
}

func (m *Model) TickAmount() float64 {
	return 0.0
}

func (m *Model) Prestige() {
}

func (m *Model) PrestigeAmount() float64 {
	return 10
}

func (m *Model) Upgrades() []upgrades.Upgrade {
	return m.upgrades
}

func (m *Model) UpgradeAvailable() bool {
	for _, upgrade := range m.upgrades {
		if upgrade.GetModel().Unlocked && !upgrade.GetModel().Enabled {
			return true
		}
	}
	return false
}

func (m Model) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func Fetch(id int, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, &ctx)
	return &layerModel
}
