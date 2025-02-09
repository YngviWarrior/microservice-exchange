package controllerdto

type InputListExchange struct{}

type OutputListExchange struct {
	Exchange int64  `json:"exchange"`
	Name     string `json:"name"`
}
