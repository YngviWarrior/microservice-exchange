package repositorydto

type InputStrategyDto struct{}
type OutputStrategyDto struct {
	Strategy uint64
	Modality uint64
	Name     string
	Enabled  bool
}
