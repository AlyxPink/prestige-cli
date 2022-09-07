package upgrades

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
)

func ListUpgrades(upgrades []Upgrade) string {
	s := strings.Builder{}
	for _, upgrade := range upgrades {
		s.WriteString(
			fmt.Sprintln(
				fmt.Sprintln(styles.MainText.Copy().Bold(true).Render(upgrade.Name)),
				fmt.Sprintln(styles.MainText.Copy().Render(upgrade.Description)),
				fmt.Sprintln(),
				fmt.Sprint(styles.MainText.Render(fmt.Sprintf("Cost: %f", upgrade.Cost))),
			),
		)
	}
	return s.String()
}
