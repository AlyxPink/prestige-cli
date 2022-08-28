package ui

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/tier_1/prestige_points"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/tier_2/boosters"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/tier_2/generators"
	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	if m.ctx.Config == nil {
		return fmt.Sprintln("Reading configuration...")
	}

	s := strings.Builder{}
	mainContent := ""

	if m.currLayer != nil {
		mainContent = lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.tiersList(),
			lipgloss.JoinVertical(
				lipgloss.Top,
				m.gameGoal(),
				m.getCurrLayer().View(),
			),
		)
	} else {
		mainContent = fmt.Sprintln("No layers found")
	}
	s.WriteString(mainContent)
	s.WriteString("\n")
	return s.String()
}

func (m *Model) fetchLayers() []layers.Layer {
	layers := []layers.Layer{
		prestige_points.Fetch(0, m.ctx),
		boosters.Fetch(1, m.ctx),
		generators.Fetch(2, m.ctx),
	}
	return layers
}

func (m *Model) setLayers(newLayers []layers.Layer) {
	m.layers = newLayers
}

func (m *Model) setCurrentLayer(layer layers.Layer) {
	m.currLayer = m.getCurrLayer()
	m.currLayerId = layer.Id()
}

func (m Model) tiersList() string {
	s := strings.Builder{}
	for _, layer := range m.fetchLayers() {
		titleStyle := styles.TierDefault
		if layer.Id() == m.currLayerId {
			titleStyle = styles.TierEnabled
		}
		s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Render(
			lipgloss.JoinHorizontal(lipgloss.Left,
				styles.PrestigeAvailable.Render(styles.DefaultGlyphs.PrestigeStatus),
				styles.UpgradeAvailable.Render(styles.DefaultGlyphs.UpgradeStatus),
				titleStyle.Copy().Render(
					fmt.Sprintf(
						"Tier: %d %s", layer.Tier(),
						layer.Name(),
					),
				),
			),
		)))
	}
	return lipgloss.NewStyle().
		Width((m.ctx.ScreenWidth / 12) * 2).
		Render(s.String())
}

func (m Model) gameGoal() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("Reach e3.140e16 points to beat the game!")))
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("You have 25,348 points! (19.25/sec)")))
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width((m.ctx.ScreenWidth / 12) * 10).
		Render(s.String())
}
