package usecase

import (
	"errors"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseUpdateAveragePrice struct {
	AvgRepo mysql.AveragePriceRepositoryInterface
}

type UsecaseUpdateAveragePriceInteface interface {
	UpdateAveragePrice(in *usecasedto.InputUpdateAveragePriceDto) (out *usecasedto.OutputUpdateAveragePriceDto, err error)
}

func NewUsecaseUpdateAveragePrice(avgRepo mysql.AveragePriceRepositoryInterface) UsecaseUpdateAveragePriceInteface {
	return &usecaseUpdateAveragePrice{
		AvgRepo: avgRepo,
	}
}

func (u *usecaseUpdateAveragePrice) UpdateAveragePrice(in *usecasedto.InputUpdateAveragePriceDto) (out *usecasedto.OutputUpdateAveragePriceDto, err error) {
	if !u.AvgRepo.Update(repositorydto.InputAveragePriceDto{
		Parity:               in.Parity,
		Exchange:             in.Exchange,
		Day:                  in.Day,
		DayUpdateTimestamp:   in.DayUpdateTimestamp,
		Week:                 in.Week,
		WeekUpdateTimestamp:  in.WeekUpdateTimestamp,
		Month:                in.Month,
		MonthUpdateTimestamp: in.MonthUpdateTimestamp,
		// Sma:                  in.Sma,
		// SmaUpdateTimestamp:   in.SmaUpdateTimestamp,
		DayRoc:   in.DayRoc,
		WeekRoc:  in.WeekRoc,
		MonthRoc: in.MonthRoc,
		// SmaRoc:               in.SmaRoc,
	}) {
		err = errors.New("error update average price")
		return
	}

	return
}
