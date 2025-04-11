package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseAvgPrices struct {
	CoinRepo mysql.CoinRepositoryInterface
}

type UseCaseAvgPricesInterface interface {
	ListAvgPrices(*usecasedto.InputAvgPricesDto) ([]*usecasedto.OutputAvgPricesDto, error)
}

func NewAvgPricesUseCase(coinRepo mysql.CoinRepositoryInterface) UseCaseAvgPricesInterface {
	return &usecaseAvgPrices{
		CoinRepo: coinRepo,
	}
}

func (u *usecaseAvgPrices) ListAvgPrices(in *usecasedto.InputAvgPricesDto) (out []*usecasedto.OutputAvgPricesDto, err error) {
	list := u.CoinRepo.FindAvg(in.From, in.To)

	out = []*usecasedto.OutputAvgPricesDto{}
	for _, v := range list {
		var o usecasedto.OutputAvgPricesDto

		o.Parity = v.Parity
		o.Exchange = v.Exchange
		o.Mts = v.Mts
		o.Open = v.Open
		o.Close = v.Close
		o.High = v.High
		o.Low = v.Low
		o.Volume = v.Volume
		o.Roc = v.Roc

		out = append(out, &o)
	}

	return
}
