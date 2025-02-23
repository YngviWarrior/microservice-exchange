package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) UpdateWallet(ctx context.Context, in *pb.UpdateWalletRequest) (out *pb.UpdateWalletResponse, err error) {
	walletRepo := mysql.NewWalletRepository(g.Db)
	usecase := usecase.NewUpdateWalletUseCase(walletRepo)

	out = &pb.UpdateWalletResponse{
		Wallet: []*pb.Wallet{},
	}
	for _, in := range in.Wallet {
		wallet, err := usecase.UpdateWallet([]*usecasedto.InputUpdateWalletDto{
			{
				Wallet:   in.GetWallet(),
				Coin:     in.GetCoin(),
				Exchange: in.GetExchange(),
				Amount:   in.GetAmount(),
			},
		})
		if err != nil {
			break
		}

		for _, w := range wallet {
			out.Wallet = append(out.Wallet, &pb.Wallet{
				Wallet:   w.Wallet,
				Exchange: w.Exchange,
				Coin:     w.Coin,
				Amount:   w.Amount,
			})
		}
	}

	return
}

func (g *grpcServer) CreateWallet(ctx context.Context, in *pb.CreateWalletRequest) (out *pb.CreateWalletResponse, err error) {
	walletRepo := mysql.NewWalletRepository(g.Db)
	usecase := usecase.NewCreateWalletUseCase(walletRepo)

	wallet, err := usecase.CreateWallet(&usecasedto.InputCreateWalletDto{
		Coin:     in.GetCoin(),
		Exchange: in.GetExchange(),
		Amount:   in.GetAmount(),
	})
	if err != nil {
		return
	}

	out = &pb.CreateWalletResponse{
		Wallet: &pb.Wallet{
			Wallet:   wallet.Wallet,
			Exchange: wallet.Exchange,
			Coin:     wallet.Coin,
			Amount:   wallet.Amount,
		},
	}

	return
}

func (g *grpcServer) ListWalletWithCoin(ctx context.Context, in *pb.ListWalletWithCoinRequest) (out *pb.ListWalletWithCoinResponse, err error) {
	walletRepo := mysql.NewWalletRepository(g.Db)
	usecase := usecase.NewListWalletWithCoinInterface(walletRepo)

	wallets, err := usecase.ListWalletWithCoin(&usecasedto.InputWalletWithCoinDto{
		Exchange: in.GetExchange(),
	})
	if err != nil {
		return
	}

	out = &pb.ListWalletWithCoinResponse{
		Wallet: []*pb.WalletWithCoin{},
	}

	for _, v := range wallets {
		var o pb.WalletWithCoin

		o.Wallet = v.Wallet
		o.Exchange = v.Exchange
		o.Coin = v.Coin
		o.Amount = v.Amount
		o.Symbol = v.Symbol
		o.Active = v.Active

		out.Wallet = append(out.Wallet, &o)
	}

	return
}

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
