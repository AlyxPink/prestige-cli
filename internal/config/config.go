package config

type LayerName string

type LayerConfig struct {
	Title string
	Name  LayerName
}

type Defaults struct {
	Layer LayerName
}

type Keybinding struct {
	Key     string
	Command string
}

type Config struct {
	Layers   []LayerConfig
	Defaults Defaults
}

func GetDefaultConfig() Config {
	return Config{
		Layers:   []LayerConfig{},
		Defaults: Defaults{},
	}
}
