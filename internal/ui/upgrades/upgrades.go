package upgrades

type Upgrade struct {
	Name        string
	Description string
	Amount      uint
	Cost        float64
	Enabled     bool
	Unlocked    bool
}

func (upgrade *Upgrade) Buy() {
	upgrade.Enabled = true
}
