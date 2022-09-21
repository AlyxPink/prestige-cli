package upgrades

type Model struct {
	Name        string
	Description string
	Amount      uint
	Cost        float64
	Enabled     bool
	Unlocked    bool
}

type Upgrade interface {
	Tick()
	Buy()
	GetModel() *Model
}
