package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
)

func (g *grpcServer) ListExchange(ctx context.Context, in *pb.ListExchangeRequest) (out *pb.ExchangeResponse, err error) {
	exchangeRepo := mysql.NewExchangeRepository(g.Db)
	useCases := usecase.NewListExchangeUsecase(exchangeRepo)

	response, err := useCases.ListExchange()
	if err != nil {
		return
	}

	out = &pb.ExchangeResponse{
		Exchange: []*pb.Exchange{},
	}

	for _, v := range response {
		var obj pb.Exchange

		obj.Exchange = uint64(v.Exchange)
		obj.Name = v.Name

		out.Exchange = append(out.Exchange, &obj)
	}

	return
}
