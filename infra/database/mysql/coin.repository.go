package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type coinRepository struct {
	Db database.DatabaseInterface
}

type CoinRepositoryInterface interface {
	List() (list []*repositorydto.OutputCoinDto)
}

func NewCoinRepository(db database.DatabaseInterface) CoinRepositoryInterface {
	return &coinRepository{
		Db: db,
	}
}

func (c *coinRepository) List() (list []*repositorydto.OutputCoinDto) {
	tx, err := c.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Panicln("CRL 00: ", err)
		return
	}

	rows, err := tx.Query(`SELECT coin, symbol, active FROM coin`)

	if err != nil {
		log.Panicln("CRL 01: ", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var c repositorydto.OutputCoinDto

		err := rows.Scan(&c.Coin, &c.Symbol, &c.Active)

		if err != nil {
			log.Panicln("CRL 02: ", err)
			return
		}

		list = append(list, &c)
	}

	return
}
