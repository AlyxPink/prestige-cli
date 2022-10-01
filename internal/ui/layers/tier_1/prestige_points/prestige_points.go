package prestige_points

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/tier_1/prestige_points/upgrades/begin"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/tier_1/prestige_points/upgrades/more_prestige"
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
		more_prestige.Fetch(pp.layer, points),
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
	if pp.PrestigeAmount() < 1 {
		return
	}
	pp.layer.Amount = pp.layer.Amount + pp.PrestigeAmount()
	pp.Points.Amount = 0
}

func (pp *PrestigePoints) PrestigeAmount() float64 {
	if pp.Points.Amount < pp.PrestigeRequirement() {
		return 0
	}
	gain := pp.Points.Amount / pp.PrestigeRequirement()
	gain = math.Pow(gain, 0.5)
	gain = gain * pp.GainMult()
	gain = math.Pow(gain, pp.GainExp())
	return gain
}

func (pp *PrestigePoints) PrestigeRequirement() float64 {
	return 10
}

func (pp *PrestigePoints) GainMult() float64 {
	mult := 1.0
	if pp.Upgrades()[3].GetModel().Unlocked { // If "more_prestige" upgrade unlocked
		mult = mult * 1.8
	}
	return mult
}

func (pp *PrestigePoints) GainExp() float64 {
	return 1
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
