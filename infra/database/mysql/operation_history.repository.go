package mysql

import (
	"database/sql"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type operationHistoryRepository struct {
	Db database.DatabaseInterface
}

type OperationHistoryRepositoryInterface interface {
	GetLastBuyRegisterByOperation(operation uint64) (coinQuantity, fee float64)
	Create(*repositorydto.InputOperationHistoryDto) bool
}

func NewOperationHistoryRepository(db database.DatabaseInterface) OperationHistoryRepositoryInterface {
	return &operationHistoryRepository{
		Db: db,
	}
}

func (t *operationHistoryRepository) Create(p *repositorydto.InputOperationHistoryDto) bool {
	stmt, err := t.Db.GetDatabase().Prepare(`
		INSERT INTO operation_history(operation, transaction_type, coin_price, coin_quantity, stable_price, stable_quantity, fee, operation_exchange_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		log.Panic("OHRC 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Operation, p.TransactionType, p.CoinPrice, p.CoinQuantity, p.StablePrice, p.StableQuantity, p.Fee, p.OperationExchangeId)

	if err != nil {
		log.Panic("OHRC 02: ", err)
		return false
	}

	return true
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
