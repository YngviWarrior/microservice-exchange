package usecase

import (
	"fmt"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseUpdateWallet struct {
	WalletRepo mysql.WalletRepositoryInterface
}

type usecaseUpdateWalletInterface interface {
	UpdateWallet(in []*usecasedto.InputUpdateWalletDto) (out []*usecasedto.OutputUpdateWalletDto, err error)
}

func NewUpdateWalletUseCase(walletRepo mysql.WalletRepositoryInterface) usecaseUpdateWalletInterface {
	return &usecaseUpdateWallet{
		WalletRepo: walletRepo,
	}
}

func (u *usecaseUpdateWallet) UpdateWallet(in []*usecasedto.InputUpdateWalletDto) (out []*usecasedto.OutputUpdateWalletDto, err error) {
	var values string

	for _, v := range in {
		values += fmt.Sprintf("(%v, %v, %v, %v),", v.Wallet, v.Exchange, v.Coin, v.Amount)
	}
	values = values[:len(values)-1]

	if !u.WalletRepo.UpdateAmount(values) {
		log.Println("UCW 00: ", err)
		return
	}

	wallets := u.WalletRepo.ListWithCoin(2)
	if err != nil {
		log.Println("UCW 01: ", err)
		return
	}

	for _, v := range wallets {
		out = append(out, &usecasedto.OutputUpdateWalletDto{
			Wallet:   v.Wallet,
			Exchange: v.Exchange,
			Coin:     v.Coin,
			Amount:   v.Amount,
		})
	}

	return
}
