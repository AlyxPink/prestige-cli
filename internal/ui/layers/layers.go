package layers

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/constants"
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Id          int
	Tier        int
	Ctx         *context.ProgramContext
	Name        string
	Amount      float64
	AmountTotal float64
	AmountBest  float64
	dimensions  constants.Dimensions
}

type Layer interface {
	Id() int
	Tier() int
	Name() string
	Tick()
	prestigable
	Upgrades() []upgrades.Upgrade
	Update(msg tea.Msg) (Layer, tea.Cmd)
	UpdateProgramContext(ctx *context.ProgramContext)
	View() string
}

type prestigable interface {
	Prestige()
	PrestigeAmount() float64
}

type LayerMsg interface {
	GetLayerId() int
}

type LayerTickMsg struct {
	LayerId         int
	InternalTickMsg tea.Msg
	Name            string
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

func (m *Model) SetDimensions(dimensions constants.Dimensions) {
	m.dimensions = dimensions
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.Ctx = ctx
	m.SetDimensions(m.GetDimensions())
}
