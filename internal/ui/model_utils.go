package ui

import "github.com/AlyxPink/prestige-cli/internal/ui/layer"

func (m *Model) getLayerAt(id int) layer.Layer {
	layers := m.layers
	if len(layers) <= id {
		return nil
	}
	return layers[id]
}

func (m *Model) getPrevLayerId() int {
	layers := m.layers
	m.currLayerId = (m.currLayerId - 1) % len(layers)
	if m.currLayerId < 0 {
		m.currLayerId += len(layers)
	}

	return m.currLayerId
}

func (m *Model) getNextLayerId() int {
	return (m.currLayerId + 1) % len(m.layers)
}
