package generators

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	tea "github.com/charmbracelet/bubbletea"
)

type Generators struct {
	layer *layers.Model
}

func NewModel(id int, ctx *context.ProgramContext) Generators {
	m := Generators{
		layer: &layers.Model{
			Id:   id,
			Tier: 2,
			Ctx:  ctx,
			Name: "Generators",
		},
	}

	return m
}

func (m *Generators) Id() int {
	return m.layer.Id
}

func (m *Generators) Name() string {
	return m.layer.Name
}

func (m *Generators) Tier() int {
	return m.layer.Tier
}

func (m *Generators) UpdateProgramContext(ctx *context.ProgramContext) {
	m.layer.UpdateProgramContext(ctx)
}

func (m *Generators) Tick() {
}

func (m *Generators) Prestige() {
}

func (m *Generators) NextPrestigeAt() float64 {
	return 10
}

func (m Generators) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func Fetch(id int, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, &ctx)
	return &layerModel
}
