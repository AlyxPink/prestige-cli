package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	warning   = lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F57DA9"}

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), true).
			BorderForeground(highlight).
			Padding(1).
			Align(lipgloss.Center)

	boxStyleEnabled = boxStyle.Copy().
			BorderForeground(special)
)

type Model struct {
	Width  int
	Height int
}

type tickMsg time.Time

func main() {
	p := tea.NewProgram(Model{}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
		return m, nil

	case tickMsg:
		return m, tick()

	}

	return m, nil
}

func (m Model) View() string {
	view := strings.Builder{}
	view.WriteString(lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.tiersList(),
		lipgloss.JoinVertical(
			lipgloss.Top,
			m.gameGoal(),
			m.stats(),
			m.prestige(),
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				m.milestones(),
				m.upgrades(),
			),
		),
	))

	return fmt.Sprintln(view.String())
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) tiersList() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Bold(true).Underline(true).Render("Tier I")))
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Render("Prestige Points")))
	s.WriteString(fmt.Sprintln())
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Bold(true).Underline(true).Render("Tier II")))
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Render("Booster")))
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Render("Generator")))
	return lipgloss.NewStyle().
		Width((m.Width / 12) * 2).
		Render(s.String())
}

func (m Model) gameGoal() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Render("Reach e3.140e16 points to beat the game!")))
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Render("You have 25,348 points!")))
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Render("(19.25/sec)")))
	return lipgloss.NewStyle().
		Width((m.Width / 12) * 10).
		Align(lipgloss.Center).
		Render(s.String())
}

func (m Model) stats() string {
	s1 := strings.Builder{}
	s1.WriteString(fmt.Sprintln(lipgloss.NewStyle().Underline(true).Render("You have:")))
	s1.WriteString(fmt.Sprintln(lipgloss.NewStyle().MarginLeft(2).Render("2 generators, generating 3.00 Generator Power/sec")))
	s1.WriteString(fmt.Sprintln(lipgloss.NewStyle().MarginLeft(2).Render("5,703 Generator Power")))
	s1.WriteString(fmt.Sprintln(lipgloss.NewStyle().MarginLeft(4).Render("boosting Point generation by 17.87x")))
	s2 := strings.Builder{}
	s2.WriteString(fmt.Sprintln())
	s2.WriteString(fmt.Sprintln(lipgloss.NewStyle().Render("Your best generators is 18")))
	s2.WriteString(fmt.Sprintln(lipgloss.NewStyle().Render("Total of 64 generators")))
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width((m.Width/12)*5).Align(lipgloss.Left).Render(s1.String()),
		lipgloss.NewStyle().Width((m.Width/12)*5).Align(lipgloss.Left).Render(s2.String()),
	)
}

func (m Model) prestige() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Underline(true).Render("Prestige")))

	button := strings.Builder{}
	button.WriteString(fmt.Sprintln(boxStyle.Copy().Width((m.Width / 12) * 4).Render(
		fmt.Sprint(
			fmt.Sprintln("Reset for +1 generators"),
			fmt.Sprint("Require: 25,348 / 40,000 points"),
		),
	)))

	return lipgloss.NewStyle().
		Width((m.Width / 12) * 10).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				button.String(),
			),
		)
}

func (m Model) milestones() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Underline(true).Render("Milestones")))

	milestones := strings.Builder{}
	milestones.WriteString(fmt.Sprintln(boxStyleEnabled.Copy().Width((m.Width / 12) * 3).Align(lipgloss.Left).Render(
		fmt.Sprint(
			fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("8 generators")),
			fmt.Sprint(lipgloss.NewStyle().Render("Keep prestige points on reset")),
		),
	)))
	milestones.WriteString(fmt.Sprintln(boxStyle.Copy().Width((m.Width / 12) * 3).Align(lipgloss.Left).Render(
		fmt.Sprint(
			fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("10 generators")),
			fmt.Sprint(lipgloss.NewStyle().Render("You gain 100% prestige points every second")),
		),
	)))

	return lipgloss.NewStyle().
		Width((m.Width / 12) * 5).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				milestones.String(),
			),
		)
}

func (m Model) upgrades() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintln(lipgloss.NewStyle().Underline(true).Render("Upgrades")))

	upgrades := lipgloss.JoinHorizontal(
		lipgloss.Top,
		fmt.Sprint(boxStyleEnabled.Copy().Width(20).Height(8).Copy().Align(lipgloss.Left).Render(
			fmt.Sprint(
				fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("GP Combo")),
				fmt.Sprintln(lipgloss.NewStyle().Render("Best Generators boost Prestige Point gain.")),
				fmt.Sprintln(lipgloss.NewStyle().Render("Currently: 3.00x")),
				fmt.Sprintln(),
				fmt.Sprint(lipgloss.NewStyle().Render("Cost: 3 generators")),
			),
		)),
		fmt.Sprint(boxStyleEnabled.Copy().Width(20).Height(8).Copy().Align(lipgloss.Left).Render(
			fmt.Sprint(
				fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("GP Combo")),
				fmt.Sprintln(lipgloss.NewStyle().Render("Best Generators boost Prestige Point gain.")),
				fmt.Sprintln(lipgloss.NewStyle().Render("Currently: 3.00x")),
				fmt.Sprintln(),
				fmt.Sprint(lipgloss.NewStyle().Render("Cost: 3 generators")),
			),
		)),
		fmt.Sprint(boxStyle.Copy().Width(20).Height(8).Copy().Align(lipgloss.Left).Render(
			fmt.Sprint(
				fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("GP Combo")),
				fmt.Sprintln(lipgloss.NewStyle().Render("Best Generators boost Prestige Point gain.")),
				fmt.Sprintln(lipgloss.NewStyle().Render("Currently: 3.00x")),
				fmt.Sprintln(),
				fmt.Sprint(lipgloss.NewStyle().Render("Cost: 3 generators")),
			),
		)),
		fmt.Sprint(boxStyle.Copy().Width(20).Height(8).Copy().Align(lipgloss.Left).Render(
			fmt.Sprint(
				fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("GP Combo")),
				fmt.Sprintln(lipgloss.NewStyle().Render("Best Generators boost Prestige Point gain.")),
				fmt.Sprintln(lipgloss.NewStyle().Render("Currently: 3.00x")),
				fmt.Sprintln(),
				fmt.Sprint(lipgloss.NewStyle().Render("Cost: 3 generators")),
			),
		)),
	)

	return lipgloss.NewStyle().
		Width((m.Width / 12) * 5).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				s.String(),
				upgrades,
			),
		)
}
