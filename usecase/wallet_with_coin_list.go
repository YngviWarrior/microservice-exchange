package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListWalletWithCoin struct {
	WalletRepo mysql.WalletRepositoryInterface
}

type usecaseListWalletWithCoinInterface interface {
	ListWalletWithCoin(in *usecasedto.InputWalletWithCoinDto) (out []*usecasedto.OutputWalletWithCoinDto, err error)
}

func NewListWalletWithCoinInterface(walletRepo mysql.WalletRepositoryInterface) usecaseListWalletWithCoinInterface {
	return &usecaseListWalletWithCoin{
		WalletRepo: walletRepo,
	}
}

func (u *usecaseListWalletWithCoin) ListWalletWithCoin(in *usecasedto.InputWalletWithCoinDto) (out []*usecasedto.OutputWalletWithCoinDto, err error) {
	list := u.WalletRepo.ListWithCoin(in.Exchange)

	out = []*usecasedto.OutputWalletWithCoinDto{}
	for _, v := range list {
		var o usecasedto.OutputWalletWithCoinDto

		o.Wallet = v.Wallet
		o.Exchange = v.Exchange
		o.Coin = v.Coin
		o.Amount = v.Amount
		o.Symbol = v.Symbol
		o.Active = v.Active

		out = append(out, &o)
	}

	return
}
