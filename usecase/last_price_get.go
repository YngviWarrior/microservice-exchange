package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseGetLastPrice struct {
	CandleRepo mysql.CandleRepositoryInterface
}

type UseCaseGetLastPriceInterface interface {
	GetLastPrice(*usecasedto.InputGetLastPriceDto) (*usecasedto.OutputGetLastPriceDto, error)
}

func NewGetLastPriceUseCase(candleRepo mysql.CandleRepositoryInterface) UseCaseGetLastPriceInterface {
	return &usecaseGetLastPrice{
		CandleRepo: candleRepo,
	}
}

func (u *usecaseGetLastPrice) GetLastPrice(in *usecasedto.InputGetLastPriceDto) (out *usecasedto.OutputGetLastPriceDto, err error) {
	lastPrice := u.CandleRepo.GetLastPrice(in.Parity, in.Exchange)

	out = &usecasedto.OutputGetLastPriceDto{}
	out.Close = lastPrice

	return
}
