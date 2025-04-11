package mysql

import (
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
	FindAvg(from, to uint64) (list []*repositorydto.OutputCandleDto)
}

func NewCoinRepository(db database.DatabaseInterface) CoinRepositoryInterface {
	return &coinRepository{
		Db: db,
	}
}

func (t *coinRepository) FindAvg(from, to uint64) (list []*repositorydto.OutputCandleDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT parity, exchange, AVG(close), (((MIN(close) - MAX(close)) / MAX(close)) * 100) as roc
		FROM candle 
		WHERE mts BETWEEN ? AND ?
		GROUP BY parity, exchange
	`)

	if err != nil {
		log.Panicln("CRFA 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(from, to)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRFA 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var c repositorydto.OutputCandleDto

		err := res.Scan(&c.Parity, &c.Exchange, &c.Close, &c.Roc)

		if err != nil {
			log.Panic("CRFA 03: ", err)
		}

		list = append(list, &c)
	}

	return
}

func (t *coinRepository) List() (list []*repositorydto.OutputCoinDto) {
	rows, err := t.Db.GetDatabase().Query(`SELECT coin, symbol, active FROM coin`)

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
