package usecase

import (
	"errors"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseCreateCandles struct {
	CandleRepo mysql.CandleRepositoryInterface
}

type UsecaseCreateCandlesInterface interface {
	CreateCandles(*usecasedto.InputCreateCandlesDto) (*usecasedto.OutputCreateCandlesDto, error)
}

func NewCreateCandles(candleRepo mysql.CandleRepositoryInterface) UsecaseCreateCandlesInterface {
	return &usecaseCreateCandles{
		CandleRepo: candleRepo,
	}
}

func (u *usecaseCreateCandles) CreateCandles(in *usecasedto.InputCreateCandlesDto) (out *usecasedto.OutputCreateCandlesDto, err error) {
	var list []*repositorydto.InputCandleDto
	for _, v := range in.Candles {
		var o repositorydto.InputCandleDto
		o.Parity = v.Parity
		o.Exchange = v.Exchange
		o.Mts = v.Mts
		o.Open = v.Open
		o.Close = v.Close
		o.High = v.High
		o.Low = v.Low
		o.Volume = v.Volume

		list = append(list, &o)
	}

	if !u.CandleRepo.CreateList(list) {
		err = errors.New("error creating candles")
		return
	}

	return
}
