package upgrades

import (
	"fmt"

	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (upgrade *Upgrade) ViewUpgrade() string {
	return fmt.Sprintln(
		fmt.Sprintln(styles.MainText.Copy().Bold(true).Render(upgrade.Name)),
		fmt.Sprintln(styles.MainText.Copy().Render(upgrade.Description)),
		fmt.Sprintln(),
		fmt.Sprint(styles.MainText.Render(fmt.Sprintf("Cost: %.2f", upgrade.Cost))),
	)
}

func ListUpgrades(upgrades []Upgrade) []string {
	s := make([]string, len(upgrades))
	for _, upgrade := range upgrades {
		if upgrade.Unlocked == false {
			continue
		}
		block := styles.UpgradeBoxAvailable.Copy().Align(lipgloss.Left).Render(
			upgrade.ViewUpgrade(),
		)
		s = append(s, block)
	}
	return s
}
