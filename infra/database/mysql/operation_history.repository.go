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
	GetLastBuyRegisterByOperation(operation uint64) (coinQuantity, fee string, status uint64)
	Get(OrderExchangeId uint64) (repositorydto.OutputOperationHistoryDto, error)
	ListByOperation(operation uint64) []*repositorydto.OutputOperationHistoryDto
	Create(*repositorydto.InputOperationHistoryDto) bool
	Update(*repositorydto.InputOperationHistoryDto) bool
}

func NewOperationHistoryRepository(db database.DatabaseInterface) OperationHistoryRepositoryInterface {
	return &operationHistoryRepository{
		Db: db,
	}
}

func (t *operationHistoryRepository) ListByOperation(operation uint64) (list []*repositorydto.OutputOperationHistoryDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT operation_history, operation, transaction_type, coin_price, 
		coin_quantity, stable_price, stable_quantity, fee, 
		operation_exchange_id, operation_exchange_status
		FROM operation_history
		WHERE operation = ?
	`)

	if err != nil {
		log.Panic("OHRLBO 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(operation)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("OHRLBO 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var o repositorydto.OutputOperationHistoryDto

		err = res.Scan(&o.OperationHistory, &o.Operation, &o.TransactionType, &o.CoinPrice, &o.CoinQuantity, &o.StablePrice, &o.StableQuantity, &o.Fee, &o.OperationExchangeId, &o.OperationExchangeStatus)

		if err != nil {
			log.Panic("OHRLBO 03: ", err)
			return
		}

		list = append(list, &o)
	}

	return
}

func (t *operationHistoryRepository) Create(p *repositorydto.InputOperationHistoryDto) bool {
	stmt, err := t.Db.GetDatabase().Prepare(`
		INSERT INTO operation_history(operation, transaction_type, coin_price, coin_quantity, stable_price, stable_quantity, fee, operation_exchange_id, operation_exchange_status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		log.Panic("OHRC 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Operation, p.TransactionType, p.CoinPrice, p.CoinQuantity, p.StablePrice, p.StableQuantity, p.Fee, p.OperationExchangeId, p.OperationExchangeStatus)

	if err != nil {
		log.Panic("OHRC 02: ", err)
		return false
	}

	return true
}

func (t *operationHistoryRepository) GetLastBuyRegisterByOperation(operation uint64) (coinQuantity, fee string, status uint64) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT coin_quantity, fee, operation_exchange_status
		FROM operation_history
		WHERE operation = ? AND transaction_type = 1 
		LIMIT 1
	`)

	if err != nil {
		log.Panicln("OHRGLBRBO 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(operation).Scan(&coinQuantity, &fee, &status)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("OHRGLBRBO 02: ", err)
		return
	}

	return
}

func (t *operationHistoryRepository) Get(orderExchangeId uint64) (out repositorydto.OutputOperationHistoryDto, err error) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT operation_history, operation, transaction_type, coin_price, 
		coin_quantity, stable_price, stable_quantity, fee, 
		operation_exchange_id, operation_exchange_status
		FROM operation_history
		WHERE operation_exchange_id = ?
	`)

	if err != nil {
		log.Panicln("OHRG 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(orderExchangeId).Scan(&out.OperationHistory, &out.Operation, &out.TransactionType, &out.CoinPrice, &out.CoinQuantity, &out.StablePrice, &out.StableQuantity, &out.Fee, &out.OperationExchangeId, &out.OperationExchangeStatus)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("OHRG 02: ", err)
		return
	}

	return
}

func (t *operationHistoryRepository) Update(p *repositorydto.InputOperationHistoryDto) bool {
	stmt, err := t.Db.GetDatabase().Prepare(`
		UPDATE operation_history
		SET 
			operation_history = ?,
			operation = ?,
			transaction_type = ?,
			coin_price = ?,
			coin_quantity = ?,
			stable_price = ?,
			stable_quantity = ?,
			fee = ?,
			operation_exchange_id = ?,
			operation_exchange_status = ?
		WHERE operation_history = ?
	`)

	if err != nil {
		log.Panicln("OHRU 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.OperationHistory, p.Operation, p.TransactionType, p.CoinPrice, p.CoinQuantity, p.StablePrice, p.StableQuantity, p.Fee, p.OperationExchangeId, p.OperationExchangeStatus, p.OperationHistory)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("OHRU 02: ", err)
		return false
	}

	return true

}
