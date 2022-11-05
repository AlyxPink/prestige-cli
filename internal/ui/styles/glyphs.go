package styles

type Glyphs struct {
	PrestigeStatusUnavailable string
	PrestigeStatusAvailable   string

	UpgradeStatusUnavailable string
	UpgradeStatusAvailable   string
}

var DefaultGlyphs = Glyphs{
	PrestigeStatusUnavailable: "○",
	PrestigeStatusAvailable:   "●",

	UpgradeStatusUnavailable: "△",
	UpgradeStatusAvailable:   "▲",
}
