package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type strategyRepository struct {
	Db database.DatabaseInterface
}

type StrategyRepositoryInterface interface {
	List(repositorydto.InputStrategyDto) []*repositorydto.OutputStrategyDto
}

func NewStrategyRepository(db database.DatabaseInterface) StrategyRepositoryInterface {
	return &strategyRepository{
		Db: db,
	}
}

func (m strategyRepository) List(in repositorydto.InputStrategyDto) (out []*repositorydto.OutputStrategyDto) {
	tx, err := m.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Panicln("SRL 00 :", err)
	}

	stmt, err := tx.Prepare(`
		SELECT strategy, modality, name, enabled
		FROM strategy
	`)

	if err != nil {
		log.Panicln("SRL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Panicln("SRL 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var u repositorydto.OutputStrategyDto

		err := res.Scan(&u.Strategy, &u.Modality, &u.Name, &u.Enabled)

		if err != nil {
			log.Panicln("SRL 03: ", err)
			return
		}

		out = append(out, &u)
	}

	tx.Commit()

	return
}
