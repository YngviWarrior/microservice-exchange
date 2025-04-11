package mysql

import (
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

func (t *parityRepository) List(in *repositorydto.InputParityDto) (list []*repositorydto.OutputParityDto) {
	stmt, err := t.Db.GetDatabase().Query(`SELECT parity, symbol, active FROM parity WHERE active = 1`)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("PRL 01: ", err)
		return
	}

	defer stmt.Close()

	for stmt.Next() {
		var c repositorydto.OutputParityDto

		err := stmt.Scan(&c.Parity, &c.Symbol, &c.Active)

		if err != nil {
			log.Panicln("CRL 02: ", err)
			return
		}

		list = append(list, &c)
	}

	return
}
