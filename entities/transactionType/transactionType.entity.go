package transactiontype

type TransactionType struct {
	TransactionType int64  `json:"transaction_type"`
	Description     string `json:"description"`
	Active          bool   `json:"active"`
}
