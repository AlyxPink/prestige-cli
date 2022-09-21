package upgrades

import (
	"fmt"

	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (upgrade *Model) ViewUpgrade() string {
	available := styles.UpgradeBoxAvailable.Copy().Align(lipgloss.Left).Render(
		fmt.Sprintln(
			fmt.Sprintln(styles.MainText.Copy().Bold(true).Render(upgrade.Name)),
			fmt.Sprintln(styles.MainText.Copy().Render(upgrade.Description)),
			fmt.Sprintln(),
			fmt.Sprint(styles.MainText.Render(fmt.Sprintf("Cost: %.2f", upgrade.Cost))),
		))

	enabled := styles.UpgradeBoxEnabled.Copy().Align(lipgloss.Left).Render(
		fmt.Sprintln(
			fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render(upgrade.Name)),
			fmt.Sprintln(styles.SubtleMainText.Copy().Render(upgrade.Description)),
			fmt.Sprintln(),
			fmt.Sprint(styles.SubtleMainText.Render(fmt.Sprintf("Cost: %.2f", upgrade.Cost))),
		))

	if upgrade.Enabled {
		return enabled
	}
	return available
}

func ListUpgrades(upgrades []Model) []string {
	s := make([]string, len(upgrades))
	for _, upgrade := range upgrades {
		if upgrade.Unlocked == false {
			continue
		}
		block := upgrade.ViewUpgrade()
		s = append(s, block)
	}
	return s
}
