package prestige_points

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	tea "github.com/charmbracelet/bubbletea"
)

type PrestigePoints struct {
	Points *points.Points
	layer  *layers.Model
}

func NewModel(id int, points *points.Points, ctx *context.ProgramContext) PrestigePoints {
	m := PrestigePoints{
		Points: points,
		layer: &layers.Model{
			Id:   id,
			Tier: 1,
			Ctx:  ctx,
			Name: "Prestige Points",
		},
	}

	return m
}

func (m *PrestigePoints) Id() int {
	return m.layer.Id
}

func (m *PrestigePoints) Name() string {
	return m.layer.Name
}

func (m *PrestigePoints) Tier() int {
	return m.layer.Tier
}

func (m *PrestigePoints) UpdateProgramContext(ctx *context.ProgramContext) {
	m.layer.UpdateProgramContext(ctx)
}

func (m *PrestigePoints) Tick() {
}

func (m *PrestigePoints) Prestige() {
	if m.Points.Amount >= 10 {
		m.Points.Amount = 0
		m.layer.Count++
	}
}

func (m *PrestigePoints) NextPrestigeAt() float64 {
	return 10
}

func (m PrestigePoints) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func Fetch(id int, points *points.Points, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, points, &ctx)
	return &layerModel
}
