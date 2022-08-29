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
			m.layersList(),
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

func (m Model) layersList() string {
	currTier := 0
	s := strings.Builder{}
	for _, layer := range m.fetchLayers() {
		if currTier != layer.Tier() {
			s.WriteString(fmt.Sprintln(m.tierTitle(layer)))
			currTier = layer.Tier()
		}
		s.WriteString(m.layerTitle(layer))
	}
	return lipgloss.NewStyle().
		Width((m.ctx.ScreenWidth / 12) * 2).
		Render(s.String())
}

func (m Model) layerTitle(layer layers.Layer) string {
	titleStyle := styles.TierDefault
	if layer.Id() == m.currLayerId {
		titleStyle = styles.TierEnabled
	}
	prestigeStatus := styles.PrestigeAvailable.Render(styles.DefaultGlyphs.PrestigeStatus)
	upgradeStatus := styles.UpgradeAvailable.Render(styles.DefaultGlyphs.UpgradeStatus)
	title := titleStyle.Copy().Render(layer.Name())
	return fmt.Sprintln(lipgloss.JoinHorizontal(lipgloss.Left, prestigeStatus, upgradeStatus, title))
}

func (m Model) tierTitle(layer layers.Layer) string {
	tierTitleStyle := styles.TierTitle
	tierTitleText := fmt.Sprintf("Tier %d", layer.Tier())
	// Remove margin top for the first tier
	if layer.Tier() == 1 {
		return fmt.Sprint(tierTitleStyle.Copy().UnsetMarginTop().Render(tierTitleText))
	}
	return fmt.Sprint(tierTitleStyle.Render(tierTitleText))
}

func (m Model) gameGoal() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("Reach e3.140e16 points to beat the game!")))
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("You have %d points! (19.25/sec)", m.points))))
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width((m.ctx.ScreenWidth / 12) * 10).
		Render(s.String())
}
