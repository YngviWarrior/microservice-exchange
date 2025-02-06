package repositorydto

type InputExchangeDto struct{}
type OutputExchangeDto struct {
	Exchange int64  `json:"exchange"`
	Name     string `json:"name"`
}
