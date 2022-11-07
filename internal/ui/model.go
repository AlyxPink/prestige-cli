package ui

import (
	"strconv"
	"time"

	"github.com/VictorBersy/prestige-cli/internal/ui/constants"
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer/boosters"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer/generators"
	"github.com/VictorBersy/prestige-cli/internal/ui/layer/prestige_points"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	tickMsg time.Time
	Model   struct {
		Tick        constants.Tick
		keys        utils.KeyMap
		err         error
		currLayerId int
		currLayer   layer.Layer
		layers      []layer.Layer
		ctx         context.ProgramContext
	}
)

func NewModel() Model {
	m := Model{
		keys: utils.Keys,
		ctx:  context.ProgramContext{},
		Tick: constants.Tick{
			Duration: time.Millisecond * 10,
		},
	}

	layers := &layer.Layers{}
	layers.Points = points.Fetch()
	layers.PrestigePoints = prestige_points.Fetch(0, layers, m.ctx)
	layers.Boosters = boosters.Fetch(1, layers, m.ctx)
	layers.Generators = generators.Fetch(2, layers, m.ctx)

	m.layers = []layer.Layer{
		layers.PrestigePoints,
		layers.Boosters,
		layers.Generators,
	}
	m.currLayer = m.layers[0]
	m.currLayerId = 0

	return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.PrevLayer):
			prevLayer := m.getLayerAt(m.getPrevLayerId())
			m.setCurrentLayer(prevLayer)

		case key.Matches(msg, m.keys.NextLayer):
			nextLayer := m.getLayerAt(m.getNextLayerId())
			m.setCurrentLayer(nextLayer)

		case key.Matches(msg, key.NewBinding(
			key.WithKeys("f1", "f2", "f3", "f4", "f5", "f6", "f7", "f8", "f9", "f10", "f11", "f12"),
		)):
			keyDigit, _ := strconv.Atoi(msg.String()[1:])
			upgradeId := keyDigit - 1
			if len(m.currLayer.Model().ListUpgrades()) > upgradeId {
				m.currLayer.Model().ListUpgrades()[upgradeId].Buy()
			}

		case key.Matches(msg, m.keys.Prestige):
			m.currLayer.Prestige()

		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit
		}

	case tickMsg:
		m.tickLayers()
		cmd = tea.Batch(tickCmd(m))

	case initMsg:
		cmd = tea.Batch(tickCmd(m))

	case layer.LayerMsg:
		cmd = m.updateCurrentLayer(msg)

	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)

	case errMsg:
		m.err = msg
	}

	m.syncProgramContext()
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func tickCmd(m Model) tea.Cmd {
	return tea.Every(m.Tick.Duration, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *Model) setCurrentLayer(layer layer.Layer) {
	m.currLayer = layer
	m.currLayerId = layer.Model().Id
}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
	m.ctx.Width = msg.Width
	m.ctx.Height = msg.Height
}

func (m *Model) syncProgramContext() {
	for _, layer := range m.layers {
		layer.UpdateProgramContext(&m.ctx)
	}
}

func (m *Model) TickPerSecond() float64 {
	amount := 0.0
	for _, layer := range m.layers {
		amount = amount + layer.TickAmount()
	}
	return amount
}

func (m *Model) tickLayers() {
	for _, layer := range m.layers {
		// Tick only if layer has been unlocked
		if !layer.Unlocked() {
			return
		}
		// Tick layer
		layer.Tick()
		// Tick upgrades
		for _, upgrade := range layer.Model().Upgrades {
			if upgrade.Model().Enabled {
				upgrade.Tick()
			}
		}
		// Tick milestones
		for _, milestone := range layer.Model().Milestones {
			if milestone.Model().Enabled {
				milestone.Tick()
			}
		}
		// Tick achievements
		for _, achievement := range layer.Model().Achievements {
			if !achievement.Model().Achieved {
				achievement.Tick()
			}
		}
	}
}

func (m *Model) updateCurrentLayer(msg layer.LayerMsg) (cmd tea.Cmd) {
	var updatedLayer layer.Layer

	updatedLayer, cmd = m.layers[msg.GetLayerId()].Update(msg)
	m.layers[msg.GetLayerId()] = updatedLayer

	return cmd
}
