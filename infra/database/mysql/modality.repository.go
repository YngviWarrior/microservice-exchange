package mysql

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type modalityRepository struct {
	Db database.DatabaseInterface
}

type ModalityRepositoryInterface interface {
	List(repositorydto.InputModalityDto) []*repositorydto.OutputModalityDto
}

func NewModalityRepository(db database.DatabaseInterface) ModalityRepositoryInterface {
	return &modalityRepository{
		Db: db,
	}
}

func (t modalityRepository) List(in repositorydto.InputModalityDto) (out []*repositorydto.OutputModalityDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT modality, name, enabled
		FROM modality
	`)

	if err != nil {
		log.Panicln("MRL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Panicln("MRL 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var u repositorydto.OutputModalityDto

		err := res.Scan(&u.Modality, &u.Name, &u.Enabled)

		if err != nil {
			log.Panicln("MRL 03: ", err)
			return
		}

		out = append(out, &u)
	}

	return
}
