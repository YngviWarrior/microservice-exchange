package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseGetFirstPrice struct {
	CandleRepo mysql.CandleRepositoryInterface
}

type UseCaseGetFirstPriceInterface interface {
	GetFirstPrice(*usecasedto.InputGetFirstPriceDto) (*usecasedto.OutputGetFirstPriceDto, error)
}

func NewGetFirstPriceUseCase(candleRepo mysql.CandleRepositoryInterface) UseCaseGetFirstPriceInterface {
	return &usecaseGetFirstPrice{
		CandleRepo: candleRepo,
	}
}

func (u *usecaseGetFirstPrice) GetFirstPrice(in *usecasedto.InputGetFirstPriceDto) (out *usecasedto.OutputGetFirstPriceDto, err error) {
	candle := u.CandleRepo.FindFirstPrice(in.Parity, in.Exchange, in.From)

	out = &usecasedto.OutputGetFirstPriceDto{}
	out.Parity = candle.Parity
	out.Exchange = candle.Exchange
	out.Mts = candle.Mts
	out.Open = candle.Open
	out.Close = candle.Close
	out.High = candle.High
	out.Low = candle.Low
	out.Volume = candle.Volume
	out.Roc = candle.Roc

	return
}
