package generators

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	layer *layer.Model
}

func NewModel(id int, layers *layer.Layers, ctx *context.ProgramContext) Model {
	g := Model{
		layer: &layer.Model{
			Name: "Generators",
			Id:   id,
			Tier: 2,
			Required: map[float64]layer.Layer{
				200: layers.Points,
			},
			Layers: layers,
			Ctx:    ctx,
		},
	}

	return g
}

func (m *Model) Model() *layer.Model {
	return m.layer
}

func (m *Model) Unlocked() bool {
	unlocked := true
	for req, layer := range m.layer.Required {
		unlocked = unlocked && layer.Model().Amount > req
	}
	return unlocked
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

func (m Model) Update(msg tea.Msg) (layer.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func Fetch(id int, layers *layer.Layers, ctx context.ProgramContext) (layer layer.Layer) {
	layerModel := NewModel(id, layers, &ctx)
	return &layerModel
}
