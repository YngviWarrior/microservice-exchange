package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) ListCoin(ctx context.Context, in *pb.ListCoinRequest) (out *pb.ListCoinResponse, err error) {
	coinRepo := mysql.NewCoinRepository(g.Db)
	usecase := usecase.NewListCoinUseCase(coinRepo)

	coins, err := usecase.ListCoin(&usecasedto.InputListCoinDto{})
	if err != nil {
		return
	}

	out = &pb.ListCoinResponse{
		Coin: []*pb.Coin{},
	}

	for _, v := range coins {
		var o pb.Coin

		o.Coin = v.Coin
		o.Symbol = v.Symbol
		o.Active = v.Active
		out.Coin = append(out.Coin, &o)
	}

	return
}
