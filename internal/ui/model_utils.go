package ui

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
)

func (m *Model) getCurrLayer() layers.Layer {
	layers := m.getLayers()
	if len(layers) == 0 {
		return nil
	}
	return layers[m.currLayerId]
}

func (m *Model) getLayerAt(id int) layers.Layer {
	layers := m.getLayers()
	if len(layers) <= id {
		return nil
	}
	return layers[id]
}

func (m *Model) getPrevLayerId() int {
	layers := m.getLayers()
	m.currLayerId = (m.currLayerId - 1) % len(layers)
	if m.currLayerId < 0 {
		m.currLayerId += len(layers)
	}

	return m.currLayerId
}

func (m *Model) getNextLayerId() int {
	return (m.currLayerId + 1) % len(m.getLayers())
}
