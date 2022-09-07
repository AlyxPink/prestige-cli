package upgrades

type Upgrade struct {
	Name        string
	Description string
	Amount      uint
	Cost        float64
	Enabled     bool
	Unlocked    bool
}
