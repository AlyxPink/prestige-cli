package boosters

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer/boosters/milestone"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	layer *layer.Model
}

func NewModel(id int, layers *layer.Layers, ctx *context.ProgramContext) Model {
	m := Model{
		layer: &layer.Model{
			Name:     "Boosters",
			Id:       id,
			Tier:     2,
			Unlocked: false,
			Required: map[layer.Layer]float64{
				layers.Points: 15,
			},
			Layers: layers,
			Ctx:    ctx,
		},
	}

	m.layer.Milestones = []layer.Milestone{
		milestone.Fetch8Boosters(layers),
		milestone.Fetch15Boosters(layers),
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
	// If unlocked, then return the value
	if m.Model().Unlocked {
		return m.Model().Unlocked
	}
	// Default value to start iterating
	unlocked := true
	for layer, req := range m.layer.Required {
		unlocked = unlocked && layer.Model().Amount > req
	}
	// Save the value to the model
	m.Model().Unlocked = unlocked
	return unlocked
}

func (m *Model) Tick() {
	m.layer.Layers.Points.Model().Amount = m.layer.Layers.Points.Model().Amount + m.TickAmount()/100
}

func (m *Model) TickAmount() float64 {
	amount := 0.0
	for _, upgrade := range m.layer.Upgrades {
		if upgrade.Model().Enabled {
			amount = amount + upgrade.TickAmount()
		}
	}
	if m.layer.Amount > 0 {
		amount = amount + math.Pow(m.EffectBase(), 1)
	}
	return amount
}

func (m *Model) EffectBase() float64 {
	amount := 2.0
	amount = amount + m.layer.Amount
	amount = math.Pow(amount, 1)
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
	return mult
}

func (m *Model) GainExp() float64 {
	return 1
}

func (m *Model) Exponent() float64 {
	return 1.25
}

func (m *Model) Base() float64 {
	return 5
}

func (m Model) Update(msg tea.Msg) (layer.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func Fetch(id int, layers *layer.Layers, ctx context.ProgramContext) (layer layer.Layer) {
	layerModel := NewModel(id, layers, &ctx)
	return &layerModel
}
