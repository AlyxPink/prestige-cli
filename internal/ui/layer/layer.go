package layer

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/prestige-cli/internal/ui/constants"
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Name     string
	Id       int
	Tier     int
	Unlocked bool

	Required map[Layer]float64

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

func (m *Model) CheckUnlock() bool {
	// If unlocked, then return the value
	if m.Unlocked {
		return m.Unlocked
	}
	// Default value to start iterating
	unlocked := true
	for layer, req := range m.Required {
		unlocked = unlocked && layer.Model().Amount > req
	}
	// Save the value to the model
	m.Unlocked = unlocked
	return unlocked
}

func (m *Model) SaveBestAmount() {
	if m.Amount > m.AmountBest {
		m.AmountBest = m.Amount
	}
}

func (m *Model) ViewLocked() string {
	s := strings.Builder{}
	for layer, req := range m.Required {
		s.WriteString(fmt.Sprintf("Reach %.2f %s to unlock.", req, layer.Model().Name))
		s.WriteString("\n")
	}
	return s.String()
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
