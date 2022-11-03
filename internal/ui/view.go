package ui

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/boosters"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/generators"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/prestige_points"
	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
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
		prestige_points.Fetch(0, m.Points, m.ctx),
		boosters.Fetch(1, m.Points, m.ctx),
		generators.Fetch(2, m.ctx),
	}
	return layers
}

func (m *Model) setLayers(newLayers []layers.Layer) {
	m.layers = newLayers
}

func (m *Model) setCurrentLayer(layer layers.Layer) {
	m.currLayer = m.getCurrLayer()
	m.currLayerId = layer.Model().Id
}

func (m Model) layersList() string {
	currTier := 0
	s := strings.Builder{}
	for _, layer := range m.fetchLayers() {
		if currTier != layer.Model().Tier {
			s.WriteString(fmt.Sprintln(m.tierTitle(layer)))
			currTier = layer.Model().Tier
		}
		s.WriteString(m.layerTitle(layer))
	}
	return lipgloss.NewStyle().
		Width((m.ctx.ScreenWidth / 12) * 2).
		Render(s.String())
}

func (m Model) layerTitle(layer layers.Layer) string {
	titleStyle := styles.TierDefault
	prestigeStatus := ""
	upgradeStatus := ""
	if layer.Model().Id == m.currLayerId {
		titleStyle = styles.TierEnabled
	}
	if layer.PrestigeAmount() > 0 {
		prestigeStatus = styles.PrestigeAvailable.Render(styles.DefaultGlyphs.PrestigeStatus)
	} else {
		prestigeStatus = styles.PrestigeUnavailable.Render(styles.DefaultGlyphs.PrestigeStatus)
	}
	if layer.Model().ListUpgradeAvailable() {
		upgradeStatus = styles.UpgradeAvailable.Render(styles.DefaultGlyphs.UpgradeStatus)
	} else {
		upgradeStatus = styles.UpgradeUnavailable.Render(styles.DefaultGlyphs.UpgradeStatus)
	}
	title := titleStyle.Copy().Render(layer.Model().Name)
	return fmt.Sprintln(lipgloss.JoinHorizontal(lipgloss.Left, prestigeStatus, upgradeStatus, title))
}

func (m Model) tierTitle(layer layers.Layer) string {
	tierTitleStyle := styles.TierTitle
	tierTitleText := fmt.Sprintf("Tier %d", layer.Model().Tier)
	// Remove margin top for the first tier
	if layer.Model().Tier == 1 {
		return fmt.Sprint(tierTitleStyle.Copy().UnsetMarginTop().Render(tierTitleText))
	}
	return fmt.Sprint(tierTitleStyle.Render(tierTitleText))
}

func (m Model) gameGoal() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("Reach e3.140e16 points to beat the game!")))
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("You have %.2f points! (%.2f/sec)", m.Points.Count(), m.TickAmount()))))
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width((m.ctx.ScreenWidth / 12) * 10).
		Render(s.String())
}
