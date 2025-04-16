package repositorydto

type InputExchangeDto struct{}
type OutputExchangeDto struct {
	Exchange uint64 `json:"exchange"`
	Name     string `json:"name"`
}
