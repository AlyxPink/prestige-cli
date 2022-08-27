package ui

import (
	"time"

	"github.com/VictorBersy/prestige-cli/internal/config"
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	tickMsg time.Time
	Model   struct {
		points        float64
		keys          utils.KeyMap
		err           error
		currLayerId   int
		currLayer     layers.Layer
		currUpgradeId int
		layers        []layers.Layer
		ctx           context.ProgramContext
	}
)

func NewModel() Model {
	return Model{
		points: 10,
		keys:   utils.Keys,
		ctx:    context.ProgramContext{Config: &config.Config{}},
	}
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

		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit
		}

	case tickMsg:
		m.tickAllLayers()
		cmd = tea.Batch(tickCmd())

	case initMsg:
		m.ctx.Config = &msg.Config
		m.ctx.Layer = m.ctx.Config.Defaults.Layer
		m.syncMainContentWidth()
		m.setLayers(m.fetchLayers())
		m.setCurrentLayer(m.layers[0])
		cmd = tea.Batch(tickCmd())

	case layers.LayerMsg:
		cmd = m.updateCurrentLayer(msg)
		m.onViewedRowChanged()

	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)

	case errMsg:
		m.err = msg
	}

	m.syncProgramContext()
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *Model) getLayers() []layers.Layer {
	return m.layers
}

func (m *Model) onViewedRowChanged() {}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
	m.ctx.Width = msg.Width
	m.ctx.Height = msg.Height
}

func (m *Model) syncProgramContext() {
	for _, layer := range m.getLayers() {
		layer.UpdateProgramContext(&m.ctx)
	}
}

func (m *Model) tickAllLayers() {
	for _, layer := range m.getLayers() {
		layer.Tick()
	}
}

func (m *Model) updateCurrentLayer(msg layers.LayerMsg) (cmd tea.Cmd) {
	var updatedLayer layers.Layer

	updatedLayer, cmd = m.layers[msg.GetLayerId()].Update(msg)
	m.layers[msg.GetLayerId()] = updatedLayer

	return cmd
}

func (m *Model) syncMainContentWidth() {
	m.ctx.Width = m.ctx.ScreenWidth
}