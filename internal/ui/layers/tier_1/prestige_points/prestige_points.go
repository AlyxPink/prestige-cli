package prestige_points

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
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
		upgrades: []upgrades.Upgrade{
			{
				Name:        "Begin",
				Description: "Generate 1 Point every second.",
				Unlocked:    true,
				Cost:        1,
			},
			{
				Name:        "Prestige Boost",
				Description: "Prestige Points boost Point generation.",
				Unlocked:    true,
				Cost:        1,
			},
			{
				Name:        "Self-Synergy",
				Description: "Points boost their own generation.",
				Unlocked:    false,
				Cost:        5,
			},
			{
				Name:        "More Prestige",
				Description: "Prestige Point gain is increased by 80%.",
				Unlocked:    false,
				Cost:        20,
			},
			{
				Name:        "Upgrade Power",
				Description: "Point generation is faster based on your Prestige Upgrades bought.",
				Unlocked:    false,
				Cost:        75,
			},
			{
				Name:        "Reverse Prestige Boost",
				Description: "Prestige Point gain is boosted by your Points.",
				Unlocked:    false,
				Cost:        5_000,
			},
		},
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
}

func (pp *PrestigePoints) Prestige() {
	if pp.Points.Amount >= 10 {
		pp.Points.Amount = 0
		pp.layer.Count++
	}
}

func (pp *PrestigePoints) NextPrestigeAt() float64 {
	return 10
}

func (pp PrestigePoints) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &pp, cmd
}

func Fetch(id int, points *points.Points, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, points, &ctx)
	return &layerModel
}
