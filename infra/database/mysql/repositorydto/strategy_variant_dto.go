package repositorydto

type InputStrategyVariantDto struct {
	StrategyVariant uint64
	Strategy        uint64
	Name            string
	Enabled         bool
}

type OutputStrategyVariantDto struct {
	StrategyVariant uint64
	Strategy        uint64
	Name            string
	Enabled         bool
}
