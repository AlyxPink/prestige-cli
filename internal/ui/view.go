package ui

import (
	"fmt"
	"strings"

	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
	"github.com/AlyxPink/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	s := strings.Builder{}
	mainContent := ""
	layerContent := ""

	if m.currLayer.Model().Unlocked {
		layerContent = m.currLayer.View()
	} else {
		layerContent = m.currLayer.Model().ViewLocked()
	}

	mainContent = lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.layersList(),
		lipgloss.JoinVertical(
			lipgloss.Top,
			m.gameGoal(),
			layerContent,
		),
	)

	s.WriteString(mainContent)
	s.WriteString("\n")
	return s.String()
}

func (m Model) layersList() string {
	currTier := 0
	s := strings.Builder{}
	for _, layer := range m.layers {
		if currTier != layer.Model().Tier {
			s.WriteString(fmt.Sprintln(m.tierTitle(layer)))
			currTier = layer.Model().Tier
		}
		s.WriteString(m.layerDetails(layer))
	}
	return lipgloss.NewStyle().
		Width((m.ctx.ScreenWidth / 12) * 2).
		Render(s.String())
}

func (m Model) layerDetails(layer layer.Layer) string {
	titleStyle := styles.TierDefault
	prestigeStatus := ""
	upgradeStatus := ""
	if layer.Model().Id == m.currLayerId {
		if layer.Model().Unlocked {
			titleStyle = styles.TierEnabled
		} else {
			titleStyle = styles.TierLocked
		}
	}
	if layer.Model().Unlocked && layer.PrestigeAmount() > 0 {
		prestigeStatus = styles.PrestigeAvailable.Render(styles.DefaultGlyphs.PrestigeStatusAvailable)
	} else {
		prestigeStatus = styles.PrestigeUnavailable.Render(styles.DefaultGlyphs.PrestigeStatusUnavailable)
	}
	if layer.Model().Unlocked && layer.Model().ListUpgradeAvailable() {
		upgradeStatus = styles.UpgradeAvailable.Render(styles.DefaultGlyphs.UpgradeStatusAvailable)
	} else {
		upgradeStatus = styles.UpgradeUnavailable.Render(styles.DefaultGlyphs.UpgradeStatusUnavailable)
	}
	title := titleStyle.Copy().Render(fmt.Sprintf("%s (%.0f)", layer.Model().Name, layer.Model().Amount))
	return fmt.Sprintln(lipgloss.JoinHorizontal(lipgloss.Left, prestigeStatus, upgradeStatus, title))
}

func (m Model) tierTitle(layer layer.Layer) string {
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
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("You have %.2f points! (%.2f/sec)", m.currLayer.Model().Layers.Points.Model().Amount, m.TickPerSecond()))))
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width((m.ctx.ScreenWidth / 12) * 10).
		Render(s.String())
}
