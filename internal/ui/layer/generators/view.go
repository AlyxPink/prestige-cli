package generators

import (
	"fmt"
	"strings"

	"github.com/AlyxPink/prestige-cli/internal/ui/layer"
	"github.com/AlyxPink/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.prestige(),
			m.stats(),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.milestones(),
			m.listUpgrades(),
		),
	)
}

func (m *Model) stats() string {
	s1 := strings.Builder{}
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("You have:")))
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().MarginLeft(2).Render(fmt.Sprintf("%.0f generators", m.layer.Amount))))
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().MarginLeft(4).Render(fmt.Sprintf("boosting Point generation by %.2fx", m.TickAmount()))))
	s2 := strings.Builder{}
	s2.WriteString(fmt.Sprintln())
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render(fmt.Sprintf("Your best generators is %.0f", m.layer.AmountBest))))
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render(fmt.Sprintf("Total of %.0f generators", m.layer.AmountTotal))))
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		styles.MainText.Copy().Width((m.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s1.String()),
		styles.MainText.Copy().Width((m.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s2.String()),
	)
}

func (m *Model) prestige() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Prestige")))

	button := strings.Builder{}
	if m.PrestigeAmount() >= 1 {
		button.WriteString(fmt.Sprintln(styles.BoxStyleAvailable.Copy().Render(
			fmt.Sprint(
				fmt.Sprintf("Reset for +%.0f generators", m.PrestigeAmount()),
				fmt.Sprintln(),
				fmt.Sprintln("Require: 25,348 / 40,000 points"),
			),
		)))
	} else {
		button.WriteString(fmt.Sprintln(styles.BoxStyleUnAvailable.Copy().Render(
			fmt.Sprint(
				fmt.Sprintf("Reset for +%.0f generators", m.PrestigeAmount()),
				fmt.Sprintln(),
				fmt.Sprintln("Require: 25,348 / 40,000 points"),
			),
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

func (m *Model) milestones() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Milestones")))

	milestones := strings.Builder{}
	for _, milestone := range m.layer.Milestones {
		if milestone.Model().Reached {
			milestones.WriteString(fmt.Sprintln(styles.BoxStyleEnabled.Copy().Width((m.layer.GetDimensions().Width / 12) * 3).Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render(milestone.Model().Name)),
					fmt.Sprint(styles.SubtleMainText.Copy().Render(milestone.Model().Description)),
				),
			)))
		} else {
			milestones.WriteString(fmt.Sprintln(styles.BoxStyleUnAvailable.Copy().Width((m.layer.GetDimensions().Width / 12) * 3).Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.MainText.Copy().Bold(true).Render(milestone.Model().Name)),
					fmt.Sprint(styles.MainText.Copy().Render(milestone.Model().Description)),
				),
			)))
		}
	}

	return lipgloss.NewStyle().
		Width((m.layer.GetDimensions().Width / 12) * 4).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				milestones.String(),
			),
		)
}

func (m *Model) listUpgrades() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Upgrades")))

	for _, chunk := range layer.Chunk(m.layer.Upgrades, 4) {
		s.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Top,
			layer.List(chunk)...,
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
