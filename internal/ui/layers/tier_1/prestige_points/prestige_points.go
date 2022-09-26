package prestige_points

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/tier_1/prestige_points/upgrades/begin"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/tier_1/prestige_points/upgrades/prestige_boost"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/tier_1/prestige_points/upgrades/self_synergy"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
	"github.com/VictorBersy/prestige-cli/internal/ui/upgrades"
	tea "github.com/charmbracelet/bubbletea"
)

type PrestigePoints struct {
	Points   *points.Points
	layer    *layers.Model
	upgrades []upgrades.Upgrade
}

func NewModel(id int, points *points.Points, ctx *context.ProgramContext) PrestigePoints {
	pp := PrestigePoints{
		Points: points,
		layer: &layers.Model{
			Id:   id,
			Tier: 1,
			Ctx:  ctx,
			Name: "Prestige Points",
		},
	}

	pp.upgrades = []upgrades.Upgrade{
		begin.Fetch(pp.layer, points),
		prestige_boost.Fetch(pp.layer, points),
		self_synergy.Fetch(pp.layer, points),
	}
	return pp
}

func (pp *PrestigePoints) Id() int {
	return pp.layer.Id
}

func (pp *PrestigePoints) Name() string {
	return pp.layer.Name
}

func (pp *PrestigePoints) Tier() int {
	return pp.layer.Tier
}

func (pp *PrestigePoints) UpdateProgramContext(ctx *context.ProgramContext) {
	pp.layer.UpdateProgramContext(ctx)
}

func (pp *PrestigePoints) Tick() {
	for _, upgrade := range pp.upgrades {
		if upgrade.GetModel().Enabled {
			upgrade.Tick()
		}
	}
}

func (pp *PrestigePoints) Prestige() {
	if pp.Points.Amount >= 10 {
		pp.Points.Amount = 0
		pp.layer.Amount++
	}
}

func (pp *PrestigePoints) NextPrestigeAt() float64 {
	return 10
}

func (pp *PrestigePoints) Upgrades() []upgrades.Upgrade {
	return pp.upgrades
}

func (pp PrestigePoints) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &pp, cmd
}

func Fetch(id int, points *points.Points, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, points, &ctx)
	return &layerModel
}
