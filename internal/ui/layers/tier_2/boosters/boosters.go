package boosters

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
	tea "github.com/charmbracelet/bubbletea"
)

type Boosters struct {
	layer    *layers.Model
	upgrades []upgrades.Upgrade
}

func NewModel(id int, ctx *context.ProgramContext) Boosters {
	b := Boosters{
		layer: &layers.Model{
			Id:   id,
			Tier: 2,
			Ctx:  ctx,
			Name: "Boosters",
		},
	}

	return b
}

func (b *Boosters) Model() *layers.Model {
	return b.layer
}

func (b *Boosters) UpdateProgramContext(ctx *context.ProgramContext) {
	b.layer.UpdateProgramContext(ctx)
}

func (b *Boosters) Tick() {
}

func (b *Boosters) TickAmount() float64 {
	return 0.0
}

func (b *Boosters) Prestige() {
}

func (b *Boosters) PrestigeAmount() float64 {
	return 10
}

func (b *Boosters) Upgrades() []upgrades.Upgrade {
	return b.upgrades
}

func (b *Boosters) UpgradeAvailable() bool {
	for _, upgrade := range b.upgrades {
		if upgrade.GetModel().Unlocked && !upgrade.GetModel().Enabled {
			return true
		}
	}
	return false
}

func (m Boosters) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func (b *Boosters) View() string {
	return "Boosters"
}

func Fetch(id int, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, &ctx)
	return &layerModel
}
