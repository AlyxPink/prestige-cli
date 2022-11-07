package layer

type ModelAchievement struct {
	Name        string
	Description string
	Achieved    bool
	Layers      *Layers
}

type Achievement interface {
	Tick()
	Done() bool
	Effect()
	Unlocked() bool
	Model() *ModelAchievement
}
