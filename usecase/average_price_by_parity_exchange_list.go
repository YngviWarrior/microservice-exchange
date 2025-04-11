package usecase

import (
	"fmt"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseGetAvgPriceByParityExchange struct {
	AvgRepo mysql.AveragePriceRepositoryInterface
}

type UsecaseGetAvgPriceByParityExchangeInteface interface {
	GetAvgPriceByParityExchange(in *usecasedto.InputGetAvgPriceByParityExchangeDto) (out *usecasedto.OutputGetAvgPriceByParityExchangeDto, err error)
}

func NewUsecaseGetAvgPriceByParityExchange(avgRepo mysql.AveragePriceRepositoryInterface) UsecaseGetAvgPriceByParityExchangeInteface {
	return &usecaseGetAvgPriceByParityExchange{
		AvgRepo: avgRepo,
	}
}

func (u *usecaseGetAvgPriceByParityExchange) GetAvgPriceByParityExchange(in *usecasedto.InputGetAvgPriceByParityExchangeDto) (out *usecasedto.OutputGetAvgPriceByParityExchangeDto, err error) {
	avg := u.AvgRepo.FindByParityExchange(&repositorydto.InputAveragePriceDto{
		Parity:   in.Parity,
		Exchange: in.Exchange,
	})

	fmt.Println(avg)

	return
}
