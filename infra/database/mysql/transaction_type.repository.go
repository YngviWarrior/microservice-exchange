package mysql

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type transactionTypeRepository struct {
	Db database.DatabaseInterface
}

type TransactionTypeRepositoryInterface interface {
	List(*repositorydto.InputTransactionTypeDto) []*repositorydto.OutputTransactionTypeDto
}

func NewTransactionTypeRepository(db database.DatabaseInterface) TransactionTypeRepositoryInterface {
	return &transactionTypeRepository{
		Db: db,
	}
}

func (t *transactionTypeRepository) List(in *repositorydto.InputTransactionTypeDto) (out []*repositorydto.OutputTransactionTypeDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT transaction_type, description, active
		FROM transaction_type
	`)

	if err != nil {
		log.Panicln("TRL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Panicln("TRL 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var u repositorydto.OutputTransactionTypeDto

		err := res.Scan(&u.TransactionType, &u.Description, &u.Active)

		if err != nil {
			log.Panicln("TRL 03: ", err)
			return
		}

		out = append(out, &u)
	}

	return
}
