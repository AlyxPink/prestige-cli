package prestige_points

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m *PrestigePoints) View() string {
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
			m.upgrades(),
		),
	)
}

func (m *PrestigePoints) stats() string {
	s1 := strings.Builder{}
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("You have:")))
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().MarginLeft(2).Render(fmt.Sprintf("%f prestige points", m.layer.Count))))
	s2 := strings.Builder{}
	s2.WriteString(fmt.Sprintln())
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render("Your best prestige points is 767")))
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render("Total of 774 prestige points")))
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		styles.MainText.Copy().Width((m.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s1.String()),
		styles.MainText.Copy().Width((m.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s2.String()),
	)
}

func (m *PrestigePoints) prestige() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Prestige")))

	button := strings.Builder{}
	button.WriteString(fmt.Sprintln(styles.BoxStyleAvailable.Copy().Render(
		fmt.Sprint(
			fmt.Sprintln("Reset for +1 prestige points"),
		),
	)))

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

func (m *PrestigePoints) milestones() string {
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

func (m *PrestigePoints) upgrades() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Upgrades")))

	upgrades := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.UpgradeBoxAvailable.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.MainText.Copy().Bold(true).Render("Begin")),
					fmt.Sprintln(styles.MainText.Copy().Render("Generate 1 Point every second.")),
					fmt.Sprintln(),
					fmt.Sprint(styles.MainText.Render("Cost: 1 prestige points")),
				),
			),
			styles.UpgradeBoxUnAvailable.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.MainText.Copy().Bold(true).Render("Prestige Boost")),
					fmt.Sprintln(styles.MainText.Copy().Render("Prestige Points boost Point generation.")),
					fmt.Sprintln(),
					fmt.Sprintln(styles.MainText.Copy().Render("Currently: 0x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.MainText.Copy().Render("Cost: 1 prestige points")),
				),
			),
			styles.UpgradeBoxUnAvailable.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.MainText.Copy().Bold(true).Render("Self-Synergy")),
					fmt.Sprintln(styles.MainText.Copy().Render("Points boost their own generation.")),
					fmt.Sprintln(),
					fmt.Sprintln(styles.MainText.Copy().Render("Currently: 7.82x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.MainText.Copy().Render("Cost: 5 prestige points")),
				),
			),
			styles.UpgradeBoxUnAvailable.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.MainText.Copy().Bold(true).Render("More Prestige")),
					fmt.Sprintln(styles.MainText.Copy().Render("Prestige Point gain is increased by 80%.")),
					fmt.Sprintln(),
					fmt.Sprint(styles.MainText.Copy().Render("Cost: 20 prestige points")),
				),
			),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.UpgradeBoxUnAvailable.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.MainText.Copy().Bold(true).Render("Upgrade Power")),
					fmt.Sprintln(styles.MainText.Copy().Render("Point generation is faster based on your Prestige Upgrades bought.")),
					fmt.Sprintln(),
					fmt.Sprintln(styles.MainText.Copy().Render("Currently: 7.53x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.MainText.Copy().Render("Cost: 75 prestige points")),
				),
			),
			styles.UpgradeBoxUnAvailable.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.MainText.Copy().Bold(true).Render("Reverse Prestige Boost")),
					fmt.Sprintln(styles.MainText.Copy().Render("Prestige Point gain is boosted by your Points.")),
					fmt.Sprintln(),
					fmt.Sprintln(styles.MainText.Copy().Render("Currently: 3.35x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.MainText.Copy().Render("Cost: 5,000 prestige points")),
				),
			),
		),
	)

	return lipgloss.NewStyle().
		Width((m.layer.GetDimensions().Width / 12) * 6).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				upgrades,
			),
		)
}
