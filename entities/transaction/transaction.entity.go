package transaction

type Transaction struct {
	Transaction     int64   `json:"transaction"`
	TransactionType int64   `json:"transaction_type"`
	WalletSent      int64   `json:"wallet_sent"`
	WalletReceived  int64   `json:"wallet_received"`
	Amount          float64 `json:"amount"`
}
