package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) GetWalletWithCoin(ctx context.Context, in *pb.GetWalletWithCoinRequest) (out *pb.GetWalletWithCoinResponse, err error) {
	walletRepo := mysql.NewWalletRepository(g.Db)
	usecases := usecase.NewGetWalletWithCoinUsecase(walletRepo)

	wallet, err := usecases.GetWalletWithCoin(&usecasedto.InputWalletWithCoinDto{
		Exchange: in.GetExchange(),
		Coin:     in.GetCoin(),
	})
	if err != nil {
		return
	}

	out = &pb.GetWalletWithCoinResponse{
		Wallet: &pb.Wallet{
			Wallet:   wallet.Wallet,
			Exchange: wallet.Exchange,
			Coin:     wallet.Coin,
			Amount:   wallet.Amount,
		},
		Coin:   wallet.Coin,
		Symbol: wallet.Symbol,
	}

	return
}
