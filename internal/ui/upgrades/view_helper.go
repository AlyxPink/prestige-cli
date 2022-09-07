package upgrades

import (
	"fmt"

	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func ListUpgrades(upgrades []Upgrade) []string {
	s := make([]string, len(upgrades))
	for _, upgrade := range upgrades {
		block := styles.UpgradeBoxAvailable.Copy().Align(lipgloss.Left).Render(
			fmt.Sprintln(
				fmt.Sprintln(styles.MainText.Copy().Bold(true).Render(upgrade.Name)),
				fmt.Sprintln(styles.MainText.Copy().Render(upgrade.Description)),
				fmt.Sprintln(),
				fmt.Sprint(styles.MainText.Render(fmt.Sprintf("Cost: %f", upgrade.Cost))),
			),
		)
		s = append(s, block)
	}
	return s
}
