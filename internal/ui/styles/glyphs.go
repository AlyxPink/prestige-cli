package styles

type Glyphs struct {
	PrestigeStatus string
	UpgradeStatus  string
}

var DefaultGlyphs = Glyphs{
	PrestigeStatus: "●",
	UpgradeStatus:  "▲",
}
