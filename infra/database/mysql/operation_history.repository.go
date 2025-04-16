package mysql

import (
	"database/sql"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
)

type operationHistoryRepository struct {
	Db database.DatabaseInterface
}

type OperationHistoryRepositoryInterface interface {
	GetLastBuyRegisterByOperation(operation uint64) (coinQuantity, fee float64)
}

func NewOperationHistoryRepository(db database.DatabaseInterface) OperationHistoryRepositoryInterface {
	return &operationHistoryRepository{
		Db: db,
	}
}

func (t *operationHistoryRepository) GetLastBuyRegisterByOperation(operation uint64) (coinQuantity, fee float64) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT coin_quantity, fee
		FROM operation_history
		WHERE operation = ? AND transaction_type = 1 
		LIMIT 1
	`)

	if err != nil {
		log.Panicln("OHRGLBRBO 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(operation).Scan(&coinQuantity, &fee)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("OHRGLBRBO 02: ", err)
		return
	}

	return
}
