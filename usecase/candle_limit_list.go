package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListCandleLimit struct {
	CandleRepo mysql.CandleRepositoryInterface
}

type UsecaseListCandleLimitInterface interface {
	ListCandleLimit(*usecasedto.InputListCandleLimitDto) (*usecasedto.OutputListCandleLimitDto, error)
}

func NewListCandleLimit(candleRepo mysql.CandleRepositoryInterface) UsecaseListCandleLimitInterface {
	return &usecaseListCandleLimit{
		CandleRepo: candleRepo,
	}
}

func (u *usecaseListCandleLimit) ListCandleLimit(in *usecasedto.InputListCandleLimitDto) (out *usecasedto.OutputListCandleLimitDto, err error) {
	candles := u.CandleRepo.ListLimit(int64(in.Parity), int64(in.Exchange), int64(in.Limit))

	out = &usecasedto.OutputListCandleLimitDto{
		Candles: []*usecasedto.Candle{},
	}

	for _, v := range candles {
		out.Candles = append(out.Candles, &usecasedto.Candle{
			Parity:   uint64(v.Parity),
			Exchange: uint64(v.Exchange),
			Mts:      uint64(v.Mts),
			Open:     v.Open,
			Close:    v.Close,
			High:     v.High,
			Low:      v.Low,
			Volume:   v.Volume,
			Roc:      v.Roc,
		})
	}

	return
}
