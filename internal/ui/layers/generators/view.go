package generators

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (g *Generators) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			g.prestige(),
			g.stats(),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			g.milestones(),
			g.listUpgrades(),
		),
	)
}

func (g *Generators) stats() string {
	s1 := strings.Builder{}
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("You have:")))
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().MarginLeft(2).Render("2 generators, generating 3.00 Generator Power/sec")))
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().MarginLeft(2).Render("5,703 Generator Power")))
	s1.WriteString(fmt.Sprintln(styles.MainText.Copy().MarginLeft(4).Render("boosting Point generation by 17.87x")))
	s2 := strings.Builder{}
	s2.WriteString(fmt.Sprintln())
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render("Your best generators is 18")))
	s2.WriteString(fmt.Sprintln(styles.MainText.Copy().Render("Total of 64 generators")))
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		styles.MainText.Copy().Width((g.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s1.String()),
		styles.MainText.Copy().Width((g.layer.GetDimensions().Width/12)*4).Align(lipgloss.Left).Render(s2.String()),
	)
}

func (g *Generators) prestige() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Prestige")))

	button := strings.Builder{}
	button.WriteString(fmt.Sprintln(styles.BoxStyleAvailable.Copy().Render(
		fmt.Sprint(
			fmt.Sprintln("Reset for +1 generators"),
			fmt.Sprintln("Require: 25,348 / 40,000 points"),
		),
	)))

	return lipgloss.NewStyle().
		Width((g.layer.GetDimensions().Width / 12) * 4).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				button.String(),
			),
		)
}

func (g *Generators) milestones() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Milestones")))

	milestones := strings.Builder{}
	milestones.WriteString(fmt.Sprintln(styles.BoxStyleEnabled.Copy().Width((g.layer.GetDimensions().Width / 12) * 3).Align(lipgloss.Left).Render(
		fmt.Sprint(
			fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render("8 generators")),
			fmt.Sprint(styles.SubtleMainText.Copy().Render("Keep prestige points on reset")),
		),
	)))
	milestones.WriteString(fmt.Sprintln(styles.BoxStyleUnAvailable.Copy().Width((g.layer.GetDimensions().Width / 12) * 3).Align(lipgloss.Left).Render(
		fmt.Sprint(
			fmt.Sprintln(styles.MainText.Copy().Bold(true).Render("10 generators")),
			fmt.Sprint(styles.MainText.Copy().Render("You gain 100% prestige points every second")),
		),
	)))

	return lipgloss.NewStyle().
		Width((g.layer.GetDimensions().Width / 12) * 4).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				milestones.String(),
			),
		)
}

func (g *Generators) listUpgrades() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(styles.MainText.Copy().Bold(true).Underline(true).Render("Upgrades")))

	upgrades := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.UpgradeBoxEnabled.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render("GP Combo")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Best Generators boost Prestige Point gain.")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Currently: 3.00x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.SubtleMainText.Render("Cost: 3 generators")),
				),
			),
			styles.UpgradeBoxEnabled.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render("GP Combo")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Best Generators boost Prestige Point gain.")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Currently: 3.00x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.SubtleMainText.Copy().Render("Cost: 3 generators")),
				),
			),
			styles.UpgradeBoxEnabled.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render("GP Combo")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Best Generators boost Prestige Point gain.")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Currently: 3.00x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.SubtleMainText.Copy().Render("Cost: 3 generators")),
				),
			),
			styles.UpgradeBoxEnabled.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render("GP Combo")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Best Generators boost Prestige Point gain.")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Currently: 3.00x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.SubtleMainText.Copy().Render("Cost: 3 generators")),
				),
			),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.UpgradeBoxEnabled.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render("GP Combo")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Best Generators boost Prestige Point gain.")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Currently: 3.00x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.SubtleMainText.Copy().Render("Cost: 3 generators")),
				),
			),
			styles.UpgradeBoxEnabled.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render("GP Combo")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Best Generators boost Prestige Point gain.")),
					fmt.Sprintln(styles.SubtleMainText.Copy().Render("Currently: 3.00x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.SubtleMainText.Copy().Render("Cost: 3 generators")),
				),
			),
			styles.UpgradeBoxAvailable.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.MainText.Copy().Bold(true).Render("GP Combo")),
					fmt.Sprintln(styles.MainText.Copy().Render("Best Generators boost Prestige Point gain.")),
					fmt.Sprintln(styles.MainText.Copy().Render("Currently: 3.00x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.MainText.Copy().Render("Cost: 3 generators")),
				),
			),
			styles.UpgradeBoxUnAvailable.Copy().Align(lipgloss.Left).Render(
				fmt.Sprint(
					fmt.Sprintln(styles.MainText.Copy().Bold(true).Render("GP Combo")),
					fmt.Sprintln(styles.MainText.Copy().Render("Best Generators boost Prestige Point gain.")),
					fmt.Sprintln(styles.MainText.Copy().Render("Currently: 3.00x")),
					fmt.Sprintln(),
					fmt.Sprint(styles.MainText.Copy().Render("Cost: 3 generators")),
				),
			),
		),
	)

	return lipgloss.NewStyle().
		Width((g.layer.GetDimensions().Width / 12) * 6).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				upgrades,
			),
		)
}
