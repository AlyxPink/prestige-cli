package layer

import "github.com/charmbracelet/lipgloss"

var (
	ContainerPadding = 1
	ContainerStyle   = lipgloss.NewStyle().
				Padding(ContainerPadding)

	MainContentBlock = ContainerStyle.Copy().Border(lipgloss.NormalBorder(), true)
	SidebarBlock     = ContainerStyle.Copy().Border(lipgloss.NormalBorder(), true)
)
