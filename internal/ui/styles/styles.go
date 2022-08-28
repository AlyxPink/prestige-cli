package styles

import "github.com/charmbracelet/lipgloss"

var (
	subtle       = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	text         = lipgloss.AdaptiveColor{Light: "#383838", Dark: "#D9DCCF"}
	textDisabled = lipgloss.AdaptiveColor{Light: "#888888", Dark: "#82847c"}
	highlight    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special      = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	warning      = lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F57DA9"}

	statusCircle        = lipgloss.NewStyle().PaddingRight(1)
	PrestigeAvailable   = statusCircle.Copy().Foreground(special)
	UpgradeAvailable    = statusCircle.Copy().Foreground(highlight)
	PrestigeUnavailable = statusCircle.Copy().Foreground(textDisabled).Bold(true)
	UpgradeUnavailable  = statusCircle.Copy().Foreground(textDisabled)

	MainText          = lipgloss.NewStyle().Foreground(text)
	SubtleMainText    = lipgloss.NewStyle().Foreground(subtle)
	DisabledTextStyle = lipgloss.NewStyle().Foreground(textDisabled)

	TierDefault = MainText.Copy().
			Padding(0, 1)
	TierEnabled = MainText.Copy().
			Bold(true).
			Padding(0, 1).
			Background(highlight)

	TierTitle = MainText.Copy().
			MarginTop(1).
			Bold(true).
			Underline(true)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.ThickBorder(), true).
			BorderForeground(highlight).
			Foreground(text).
			Padding(1).
			Align(lipgloss.Center)

	BoxStyleAvailable = boxStyle.Copy().
				BorderForeground(highlight)

	BoxStyleUnAvailable = boxStyle.Copy().
				BorderForeground(warning)

	BoxStyleEnabled = boxStyle.Copy().
			BorderForeground(special).
			Background(special).
			Foreground(lipgloss.Color(subtle.Dark)).
			BorderStyle(lipgloss.Border{
			Top:         "▄",
			Bottom:      "▀",
			Left:        "▐",
			Right:       "▌",
			TopLeft:     "▗",
			TopRight:    "▖",
			BottomRight: "▘",
			BottomLeft:  "▝",
		})
)
