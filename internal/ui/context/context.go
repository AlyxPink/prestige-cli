package context

import (
	"github.com/VictorBersy/prestige-cli/internal/config"
)

type ProgramContext struct {
	ScreenHeight int
	ScreenWidth  int
	Width        int
	Height       int
	Config       *config.Config
	Layer        config.LayerName
}

func (ctx *ProgramContext) GetLayersConfig() []config.LayerConfig {
	return ctx.Config.Layers
}
