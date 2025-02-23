package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseGetCandleFirstMts struct {
	CandleRepo mysql.CandleRepositoryInterface
}

type UsecaseGetCandleFirstMtsInterface interface {
	GetCandleFirstMts(*usecasedto.InputGetCandleFirstMtsDto) (*usecasedto.OutputGetCandleFirstMtsDto, error)
}

func NewGetCandleFirstMts(candleRepo mysql.CandleRepositoryInterface) UsecaseGetCandleFirstMtsInterface {
	return &usecaseGetCandleFirstMts{
		CandleRepo: candleRepo,
	}
}

func (u *usecaseGetCandleFirstMts) GetCandleFirstMts(in *usecasedto.InputGetCandleFirstMtsDto) (out *usecasedto.OutputGetCandleFirstMtsDto, err error) {
	candle := u.CandleRepo.GetFirstMts(int64(in.Parity), int64(in.Exchange))

	out = &usecasedto.OutputGetCandleFirstMtsDto{
		Parity:   candle.Parity,
		Exchange: candle.Exchange,
		Mts:      candle.Mts,
		Open:     candle.Open,
		Close:    candle.Close,
		High:     candle.High,
		Low:      candle.Low,
		Volume:   candle.Volume,
		Roc:      candle.Roc,
	}

	return
}
