package prestige_points

import (
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/prestige_points/upgrade"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/upgrades"
	"github.com/VictorBersy/prestige-cli/internal/ui/points"
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
		upgrade.FetchBegin(pp.layer, points),
		upgrade.FetchPrestigeBoost(pp.layer, points),
		upgrade.FetchSelfSynergy(pp.layer, points),
		upgrade.FetchMorePrestige(pp.layer, points),
		upgrade.FetchReversePrestigeBoost(pp.layer, points),
	}
	return pp
}

func (pp *PrestigePoints) UpdateProgramContext(ctx *context.ProgramContext) {
	pp.layer.UpdateProgramContext(ctx)
}

func (pp *PrestigePoints) Model() *layers.Model {
	return pp.layer
}

func (pp *PrestigePoints) Tick() {
	for _, upgrade := range pp.upgrades {
		if upgrade.GetModel().Enabled {
			upgrade.Tick()
		}
	}
}

func (pp *PrestigePoints) TickAmount() float64 {
	amount := 0.0
	for _, upgrade := range pp.upgrades {
		if upgrade.GetModel().Enabled {
			amount = amount + upgrade.TickAmount()
		}
	}
	return amount
}

func (pp *PrestigePoints) Prestige() {
	if pp.PrestigeAmount() < 1 {
		return
	}
	pp.layer.Amount = pp.layer.Amount + pp.PrestigeAmount()
	pp.layer.AmountTotal = pp.layer.AmountTotal + pp.PrestigeAmount()
	// Save best score
	if pp.layer.Amount > pp.layer.AmountBest {
		pp.layer.AmountBest = pp.layer.Amount
	}
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
	if pp.Upgrades()[3].GetModel().Enabled { // If "more_prestige" upgrade enabled
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

func (pp *PrestigePoints) UpgradeAvailable() bool {
	for _, upgrade := range pp.upgrades {
		if upgrade.GetModel().Unlocked && !upgrade.GetModel().Enabled {
			return true
		}
	}
	return false
}

func (pp PrestigePoints) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &pp, cmd
}

func Fetch(id int, points *points.Points, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, points, &ctx)
	return &layerModel
}
