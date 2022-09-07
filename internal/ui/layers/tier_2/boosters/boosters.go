package boosters

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	tea "github.com/charmbracelet/bubbletea"
)

type Boosters struct {
	layer *layers.Model
}

func NewModel(id int, ctx *context.ProgramContext) Boosters {
	m := Boosters{
		layer: &layers.Model{
			Id:   id,
			Tier: 2,
			Ctx:  ctx,
			Name: "Boosters",
		},
	}

	return m
}

func (m *Boosters) Id() int {
	return m.layer.Id
}

func (m *Boosters) Name() string {
	return m.layer.Name
}

func (m *Boosters) Tier() int {
	return m.layer.Tier
}

func (m *Boosters) UpdateProgramContext(ctx *context.ProgramContext) {
	m.layer.UpdateProgramContext(ctx)
}

func (m *Boosters) Tick() {
}

func (m *Boosters) Prestige() {
}

func (m *Boosters) NextPrestigeAt() float64 {
	return 10
}

func (m Boosters) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func (m *Boosters) View() string {
	return "Boosters"
}

func Fetch(id int, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, &ctx)
	return &layerModel
}
