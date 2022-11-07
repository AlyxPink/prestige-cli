package layer

type ModelMilestone struct {
	Name                   string
	Description            string
	RequirementDescription string
	Reached                bool
	Layers                 *Layers
}

type Milestone interface {
	Tick()
	Done() bool
	Unlocked() bool
	Model() *ModelMilestone
}
