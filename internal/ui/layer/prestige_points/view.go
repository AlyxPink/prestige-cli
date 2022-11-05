package prestige_points

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/prestige-cli/internal/ui/layer/upgrades"
	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.viewPrestige(),
			m.viewStats(),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.listMilestones(),
			m.listUpgrades(),
		),
	)
}

func (m *Model) viewStats() string {
	s1 := strings.Builder{}
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("You have:")))
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().MarginLeft(2).Render(fmt.Sprintf("%.0f prestige points", m.layer.Amount))))
	s2 := strings.Builder{}
	s2.WriteString(fmt.Sprintln())
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render(fmt.Sprintf("Your best prestige points is %.0f", m.layer.AmountBest))))
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render(fmt.Sprintf("Total of %.0f prestige points", m.layer.AmountTotal))))
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		styles.MainText.Copy().Width((m.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s1.String()),
		styles.MainText.Copy().Width((m.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s2.String()),
	)
}

func (m *Model) viewPrestige() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Prestige")))

	button := strings.Builder{}
	if m.PrestigeAmount() >= 1 {
		button.WriteString(fmt.Sprintln(styles.BoxStyleAvailable.Copy().Render(
			fmt.Sprintf("Reset for +%.0f prestige points", m.PrestigeAmount()),
		)))
	} else {
		button.WriteString(fmt.Sprintln(styles.BoxStyleUnAvailable.Copy().Render(
			fmt.Sprintf("Reset for +%.0f prestige points", m.PrestigeAmount()),
		)))
	}

	return lipgloss.NewStyle().
		Width((m.layer.GetDimensions().Width / 12) * 4).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				button.String(),
			),
		)
}

func (m *Model) listMilestones() string {
	return lipgloss.NewStyle().
		Width((m.layer.GetDimensions().Width / 12) * 4).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				"",
				"",
			),
		)
}

func (m *Model) listUpgrades() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Upgrades")))

	for _, chunk := range upgrades.Chunk(m.layer.Upgrades, 4) {
		s.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Top,
			upgrades.List(chunk)...,
		))
		s.WriteRune('\n')
	}

	return lipgloss.NewStyle().
		Width((m.layer.GetDimensions().Width / 12) * 6).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
			),
		)
}
