package points

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	layer *layer.Model
}

func NewModel(id int, ctx *context.ProgramContext) Model {
	m := Model{
		layer: &layer.Model{
			Name:     "Points",
			Id:       id,
			Amount:   10,
			Unlocked: true,
			Ctx:      ctx,
		},
	}

	return m
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) { m.layer.UpdateProgramContext(ctx) }
func (m *Model) Model() *layer.Model                              { return m.layer }
func (m *Model) Unlocked() bool                                   { return true }
func (m *Model) Tick()                                            {}
func (m *Model) TickAmount() float64                              { return 0.0 }
func (m *Model) Prestige()                                        {}
func (m *Model) PrestigeAmount() float64                          { return 0.0 }
func (m *Model) PrestigeRequirement() float64                     { return 0.0 }
func (m *Model) View() string                                     { return "" }

func (m Model) Update(msg tea.Msg) (layer.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func Fetch(id int, ctx context.ProgramContext) (layer layer.Layer) {
	layerModel := NewModel(id, &ctx)
	return &layerModel
}
