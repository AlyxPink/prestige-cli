package layer

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/constants"
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Name string
	Id   int
	Tier int

	Layers *Layers

	Amount      float64
	AmountTotal float64
	AmountBest  float64

	Achievements []Achievement
	Milestones   []Milestone
	Upgrades     []Upgrade

	Ctx        *context.ProgramContext
	dimensions constants.Dimensions
}

type Layer interface {
	Tick()
	TickAmount() float64

	Prestige()
	PrestigeAmount() float64

	Unlocked() bool

	Update(msg tea.Msg) (Layer, tea.Cmd)
	View() string
	Model() *Model
	UpdateProgramContext(ctx *context.ProgramContext)
}

type Layers struct {
	Points Layer

	PrestigePoints Layer

	Boosters   Layer
	Generators Layer
}

type LayerMsg interface {
	GetLayerId() int
}

type LayerTickMsg struct {
	LayerId int
}

func (msg LayerTickMsg) GetLayerId() int {
	return msg.LayerId
}

func (m *Model) GetDimensions() constants.Dimensions {
	return constants.Dimensions{
		Width:  m.Ctx.Width,
		Height: m.Ctx.Height,
	}
}

func (m *Model) SaveBestAmount() {
	if m.Amount > m.AmountBest {
		m.AmountBest = m.Amount
	}
}

func (m *Model) ListUpgrades() []Upgrade {
	return m.Upgrades
}

func (m *Model) ListUpgradeAvailable() bool {
	for _, upgrade := range m.Upgrades {
		if upgrade.Unlocked() && !upgrade.Model().Enabled {
			return true
		}
	}
	return false
}

func (m *Model) ListUpgradeEnabled() []Upgrade {
	var upgrades_enabled []Upgrade
	for _, upgrade := range m.Upgrades {
		if upgrade.Model().Enabled {
			upgrades_enabled = append(upgrades_enabled, upgrade)
		}
	}
	return upgrades_enabled
}

func (m *Model) SetDimensions(dimensions constants.Dimensions) {
	m.dimensions = dimensions
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.Ctx = ctx
	m.SetDimensions(m.GetDimensions())
}
