package prestige_points

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
	"github.com/charmbracelet/lipgloss"
)

func (pp *PrestigePoints) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			pp.viewPrestige(),
			pp.viewStats(),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			pp.listMilestones(),
			pp.listUpgrades(),
		),
	)
}

func (pp *PrestigePoints) viewStats() string {
	s1 := strings.Builder{}
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("You have:")))
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().MarginLeft(2).Render(fmt.Sprintf("%.0f prestige points", pp.layer.Amount))))
	s2 := strings.Builder{}
	s2.WriteString(fmt.Sprintln())
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render(fmt.Sprintf("Your best prestige points is %.0f", pp.layer.AmountBest))))
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render(fmt.Sprintf("Total of %.0f prestige points", pp.layer.AmountTotal))))
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		styles.MainText.Copy().Width((pp.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s1.String()),
		styles.MainText.Copy().Width((pp.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s2.String()),
	)
}

func (pp *PrestigePoints) viewPrestige() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Prestige")))

	button := strings.Builder{}
	if pp.PrestigeAmount() >= 1 {
		button.WriteString(fmt.Sprintln(styles.BoxStyleAvailable.Copy().Render(
			fmt.Sprintf("Reset for +%.0f prestige points", pp.PrestigeAmount()),
		)))
	} else {
		button.WriteString(fmt.Sprintln(styles.BoxStyleUnAvailable.Copy().Render(
			fmt.Sprintf("Reset for +%.0f prestige points", pp.PrestigeAmount()),
		)))
	}

	return lipgloss.NewStyle().
		Width((pp.layer.GetDimensions().Width / 12) * 4).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				button.String(),
			),
		)
}

func (pp *PrestigePoints) listMilestones() string {
	return lipgloss.NewStyle().
		Width((pp.layer.GetDimensions().Width / 12) * 4).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				"",
				"",
			),
		)
}

func (pp *PrestigePoints) listUpgrades() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Upgrades")))

	u := make([]*upgrades.Model, len(pp.upgrades))
	for i, upgrade := range pp.upgrades {
		u[i] = upgrade.GetModel()
	}

	upgrades := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			upgrades.ListUpgrades(u)...,
		),
	)

	return lipgloss.NewStyle().
		Width((pp.layer.GetDimensions().Width / 12) * 6).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				upgrades,
			),
		)
}
