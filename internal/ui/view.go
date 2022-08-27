package ui

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers/prestige_points"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	if m.ctx.Config == nil {
		return fmt.Sprintln("Reading configuration...")
	}

	s := strings.Builder{}
	s.WriteString("\n")
	mainContent := ""

	if m.currLayer != nil {
		mainContent = m.getCurrLayer().View()
	} else {
		mainContent = fmt.Sprintln("No layers found")
	}
	s.WriteString(mainContent)
	s.WriteString("\n")
	return s.String()
}

func (m *Model) fetchLayers() []layers.Layer {
	layers := []layers.Layer{
		prestige_points.Fetch(m.ctx),
	}
	return layers
}

func (m *Model) setLayers(newLayers []layers.Layer) {
	m.layers = newLayers
}

func (m *Model) setCurrentLayer(layer layers.Layer) {
	m.currLayer = m.getCurrLayer()
	m.currLayerId = layer.Id()
	m.onViewedRowChanged()
}
