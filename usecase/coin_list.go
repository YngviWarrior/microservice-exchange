package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListCoin struct {
	CoinRepo mysql.CoinRepositoryInterface
}

type UseCaseListCoinInterface interface {
	ListCoin(*usecasedto.InputListCoinDto) ([]*usecasedto.OutputListCoinDto, error)
}

func NewListCoinUseCase(coinRepo mysql.CoinRepositoryInterface) UseCaseListCoinInterface {
	return &usecaseListCoin{
		CoinRepo: coinRepo,
	}
}

func (u *usecaseListCoin) ListCoin(in *usecasedto.InputListCoinDto) (out []*usecasedto.OutputListCoinDto, err error) {
	list := u.CoinRepo.List()

	out = []*usecasedto.OutputListCoinDto{}
	for _, v := range list {
		var o usecasedto.OutputListCoinDto

		o.Coin = v.Coin
		o.Symbol = v.Symbol
		o.Active = v.Active

		out = append(out, &o)
	}

	return
}
