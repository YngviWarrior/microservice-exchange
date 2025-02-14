package repositorydto

type InputStrategyDto struct{}
type OutputStrategyDto struct {
	Strategy int64
	Modality int64
	Name     string
	Enabled  bool
}
