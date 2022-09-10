package upgrades

import "fmt"

type Upgrade struct {
	Name        string
	Description string
	Amount      uint
	Cost        float64
	Enabled     bool
	Unlocked    bool
}

func (upgrade *Upgrade) Buy() {
	fmt.Printf("You just bough %s", upgrade.Name)
}
