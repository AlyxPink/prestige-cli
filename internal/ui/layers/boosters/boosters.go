package boosters

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Points *points.Points
	layer  *layers.Model
}

func NewModel(id int, points *points.Points, ctx *context.ProgramContext) Model {
	m := Model{
		Points: points,
		layer: &layers.Model{
			Id:   id,
			Tier: 2,
			Ctx:  ctx,
			Name: "Boosters",
		},
	}

	return m
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.layer.UpdateProgramContext(ctx)
}

func (m *Model) Model() *layers.Model {
	return m.layer
}

func (m *Model) Tick() {
	for _, upgrade := range m.layer.Upgrades {
		if upgrade.GetModel().Enabled {
			upgrade.Tick()
		}
	}
}

func (m *Model) TickAmount() float64 {
	amount := 0.0
	for _, upgrade := range m.layer.Upgrades {
		if upgrade.GetModel().Enabled {
			amount = amount + upgrade.TickAmount()
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
	// Save best score
	if m.layer.Amount > m.layer.AmountBest {
		m.layer.AmountBest = m.layer.Amount
	}
	m.Points.Amount = 0
}

func (m *Model) PrestigeAmount() float64 {
	if m.Points.Amount < m.PrestigeRequirement() {
		return 0
	}
	gain := m.Points.Amount / m.PrestigeRequirement()
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
	return mult
}

func (m *Model) GainExp() float64 {
	return 1
}

func (m Model) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func Fetch(id int, points *points.Points, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, points, &ctx)
	return &layerModel
}
