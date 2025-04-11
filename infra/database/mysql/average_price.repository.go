package mysql

import (
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
	FindByParityExchange(*repositorydto.InputAveragePriceDto) *repositorydto.OutputAveragePriceDto
}

func NewAveragePriceRepository(db database.DatabaseInterface) AveragePriceRepositoryInterface {
	return &averagePriceRepository{
		Db: db,
	}
}

func (t *averagePriceRepository) FindByParityExchange(in *repositorydto.InputAveragePriceDto) (a *repositorydto.OutputAveragePriceDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT parity, exchange, day, day_update_timestamp, week, week_update_timestamp, month, month_update_timestamp
		FROM average_price
		WHERE parity = ? and exchange = ?
	`)

	if err != nil {
		log.Panic("APRFBP 01: ", err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(in.Parity, in.Exchange).Scan(&a.Parity, &a.Exchange, &a.Day, &a.DayUpdateTimestamp, &a.Week, &a.WeekUpdateTimestamp, &a.Month, &a.MonthUpdateTimestamp)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("APRFBP 02: ", err)
		return
	}

	return
}

func (t *averagePriceRepository) List(in repositorydto.InputAveragePriceDto) (list []*repositorydto.OutputAveragePriceDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
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

	return
}
