package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type parityRepository struct {
	Db database.DatabaseInterface
}

type ParityRepositoryInterface interface {
	List(*repositorydto.InputParityDto) (list []*repositorydto.OutputParityDto)
}

func NewParityRepository(db database.DatabaseInterface) ParityRepositoryInterface {
	return &parityRepository{
		Db: db,
	}
}

func (p *parityRepository) List(in *repositorydto.InputParityDto) (list []*repositorydto.OutputParityDto) {
	tx, err := p.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Panicln("PRL 00 :", err)
	}

	res, err := tx.Query(`SELECT parity, symbol, active FROM parity WHERE active = 1`)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("PRL 01: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var c repositorydto.OutputParityDto

		err := res.Scan(&c.Parity, &c.Symbol, &c.Active)

		if err != nil {
			log.Panicln("CRL 02: ", err)
			return
		}

		list = append(list, &c)
	}

	return
}
