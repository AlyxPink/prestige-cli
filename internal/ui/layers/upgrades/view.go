package upgrades

import (
	"fmt"

	"github.com/VictorBersy/prestige-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View(upgrade Upgrade) string {
	var available_effect, enabled_effect string

	if upgrade.Effect() != "" {
		available_effect = fmt.Sprintln(
			fmt.Sprintln(),
			fmt.Sprintln(styles.MainText.Copy().Render(fmt.Sprint("Current: ", upgrade.Effect()))),
		)
		enabled_effect = fmt.Sprintln(
			fmt.Sprintln(),
			fmt.Sprintln(styles.SubtleMainText.Copy().Render(fmt.Sprint("Current: ", upgrade.Effect()))),
		)
	}

	available := styles.UpgradeBoxAvailable.Copy().Align(lipgloss.Left).Render(
		fmt.Sprintln(
			fmt.Sprintln(styles.MainText.Copy().Bold(true).Render(m.Name)),
			fmt.Sprintln(styles.MainText.Copy().Render(m.Description)),
			available_effect,
			fmt.Sprintln(),
			fmt.Sprint(styles.MainText.Render(fmt.Sprintf("Cost: %.2f", m.Cost))),
		))

	enabled := styles.UpgradeBoxEnabled.Copy().Align(lipgloss.Left).Render(
		fmt.Sprintln(
			fmt.Sprintln(styles.SubtleMainText.Copy().Bold(true).Render(m.Name)),
			fmt.Sprintln(styles.SubtleMainText.Copy().Render(m.Description)),
			enabled_effect,
			fmt.Sprintln(),
			fmt.Sprint(styles.SubtleMainText.Render(fmt.Sprintf("Cost: %.2f", m.Cost))),
		))

	if m.Enabled {
		return enabled
	}
	return available
}

func List(upgrades []Upgrade) []string {
	s := make([]string, len(upgrades))
	for _, upgrade := range upgrades {
		if !upgrade.Unlocked() {
			continue
		}
		block := upgrade.GetModel().View(upgrade)
		s = append(s, block)
	}
	return s
}

func Chunk(slice []Upgrade, chunkSize int) [][]Upgrade {
	var chunks [][]Upgrade
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
