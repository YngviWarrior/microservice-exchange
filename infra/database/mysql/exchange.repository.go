package mysql

import (
	"context"
	"database/sql"
	"log"

	entity "github.com/YngviWarrior/microservice-exchange/entities/exchange"
	"github.com/YngviWarrior/microservice-exchange/infra/database"
)

type exchangeRepository struct {
	Db database.DatabaseInterface
}

type ExchangeRepositoryInterface interface {
	List() (list []*entity.Exchange)
}

func NewExchangeRepository(db database.DatabaseInterface) ExchangeRepositoryInterface {
	return &exchangeRepository{
		Db: db,
	}
}

func (e *exchangeRepository) List() (list []*entity.Exchange) {
	tx, err := e.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Panicln("ERL 00 :", err)
	}

	stmt, err := tx.Prepare(`
		SELECT exchange, name
		FROM exchange
	`)

	if err != nil {
		log.Panicln("ERL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Panicln("ERL 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var u entity.Exchange

		err := res.Scan(&u.Exchange, &u.Name)

		if err != nil {
			log.Panicln("ERL 03: ", err)
			return
		}

		list = append(list, &u)
	}

	tx.Commit()

	return
}
