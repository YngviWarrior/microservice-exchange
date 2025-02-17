package usecase

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseGetWalletWithCoin struct {
	WalletRepo mysql.WalletRepositoryInterface
}

type UsecaseGetWalletWithCoinInterface interface {
	GetWalletWithCoin(in *usecasedto.InputWalletWithCoinDto) (out *usecasedto.OutputWalletWithCoinDto, err error)
}

func NewGetWalletWithCoinUsecase(walletRepo mysql.WalletRepositoryInterface) UsecaseGetWalletWithCoinInterface {
	return &usecaseGetWalletWithCoin{
		WalletRepo: walletRepo,
	}
}

func (u *usecaseGetWalletWithCoin) GetWalletWithCoin(in *usecasedto.InputWalletWithCoinDto) (out *usecasedto.OutputWalletWithCoinDto, err error) {
	response, err := u.WalletRepo.GetWithCoin(in.Exchange, in.Coin)
	if err != nil {
		log.Panicln("LTC 00: ", err)
	}

	out = &usecasedto.OutputWalletWithCoinDto{
		Wallet:   response.Wallet,
		Exchange: response.Exchange,
		Coin:     response.Coin,
		Amount:   response.Amount,
		Active:   response.Active,
	}

	return
}
