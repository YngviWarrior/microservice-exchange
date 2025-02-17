package repositorydto

type InputTransactionTypeDto struct {
	TransactionType uint64
	Description     string
	Active          bool
}

type OutputTransactionTypeDto struct {
	TransactionType uint64
	Description     string
	Active          bool
}
