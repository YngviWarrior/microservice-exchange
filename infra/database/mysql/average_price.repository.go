package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type averagePriceRepository struct {
	Db database.DatabaseInterface
}

type AveragePriceRepositoryInterface interface {
	List(repositorydto.InputAveragePriceDto) (list []*repositorydto.OutputAveragePriceDto)
}

func NewAveragePriceRepository(db database.DatabaseInterface) AveragePriceRepositoryInterface {
	return &averagePriceRepository{
		Db: db,
	}
}

func (e *averagePriceRepository) List(in repositorydto.InputAveragePriceDto) (list []*repositorydto.OutputAveragePriceDto) {
	tx, err := e.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Panicln("APRL 00 :", err)
	}

	stmt, err := tx.Prepare(`
		SELECT average_price, parity, exchange, day, day_update_timestamp, 
			week, week_update_timestamp, month, month_update_timestamp, sma, sma_update_timestamp
		FROM average_price
	`)

	if err != nil {
		log.Panicln("APRL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Panicln("APRL 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var u repositorydto.OutputAveragePriceDto

		err := res.Scan(
			&u.Parity,
			&u.Exchange,
			&u.Day,
			&u.DayUpdateTimestamp,
			&u.Week,
			&u.WeekUpdateTimestamp,
			&u.Month,
			&u.MonthUpdateTimestamp,
			&u.Sma,
			&u.SmaUpdateTimestamp,
			&u.DayRoc,
			&u.WeekRoc,
			&u.MonthRoc,
			&u.SmaRoc,
		)

		if err != nil {
			log.Panicln("APRL 03: ", err)
			return
		}

		list = append(list, &u)
	}

	tx.Commit()

	return
}
