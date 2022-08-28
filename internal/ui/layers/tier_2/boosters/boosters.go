package boosters

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	layer *layers.Model
}

func NewModel(id int, ctx *context.ProgramContext) Model {
	m := Model{
		layer: &layers.Model{
			Id:   id,
			Tier: 2,
			Ctx:  ctx,
			Name: "Boosters",
		},
	}

	return m
}

func (m *Model) Id() int {
	return m.layer.Id
}

func (m *Model) Name() string {
	return m.layer.Name
}

func (m *Model) Tier() int {
	return m.layer.Tier
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.layer.UpdateProgramContext(ctx)
}

func (m *Model) Tick() {
}

func (m Model) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func (m *Model) View() string {
	return "Boosters"
}

func Fetch(id int, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, &ctx)
	return &layerModel
}
