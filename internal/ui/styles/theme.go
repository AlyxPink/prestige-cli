package styles

import "github.com/charmbracelet/lipgloss"

var (
	indigo       = lipgloss.AdaptiveColor{Light: "#FFAFCC", Dark: "#383B5B"}
	subtleIndigo = lipgloss.AdaptiveColor{Light: "#FFC8DD", Dark: "#242347"}
)

type Theme struct {
	MainText      lipgloss.AdaptiveColor
	SubleMainText lipgloss.AdaptiveColor

	Border          lipgloss.AdaptiveColor
	SecondaryBorder lipgloss.AdaptiveColor

	SuccessText lipgloss.AdaptiveColor
	ErrorText   lipgloss.AdaptiveColor

	FaintBorder lipgloss.AdaptiveColor
	FaintText   lipgloss.AdaptiveColor

	SelectedBackground lipgloss.AdaptiveColor
	SecondaryText      lipgloss.AdaptiveColor
}

var DefaultTheme = Theme{
	MainText:      lipgloss.AdaptiveColor{Light: "#242347", Dark: "#E2E1ED"},
	SubleMainText: subtleIndigo,

	Border:          lipgloss.AdaptiveColor{Light: indigo.Light, Dark: indigo.Dark},
	SecondaryBorder: lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#39386B"},

	SuccessText: lipgloss.AdaptiveColor{Light: "#3DF294", Dark: "#06D6A0"},
	ErrorText:   lipgloss.AdaptiveColor{Light: "#F23D5C", Dark: "#EF476F"},

	FaintBorder: lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#2B2B40"},
	FaintText:   lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#3E4057"},

	SelectedBackground: lipgloss.AdaptiveColor{Light: subtleIndigo.Light, Dark: "#39386B"},
	SecondaryText:      lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#666CA6"},
}

var (
	SingleRuneWidth    = 4
	MainContentPadding = 1
)
