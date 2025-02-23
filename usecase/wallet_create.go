package usecase

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseCreateWallet struct {
	WalletRepo mysql.WalletRepositoryInterface
}

type usecaseCreateWalletInterface interface {
	CreateWallet(in *usecasedto.InputCreateWalletDto) (out *usecasedto.OutputCreateWalletDto, err error)
}

func NewCreateWalletUseCase(walletRepo mysql.WalletRepositoryInterface) usecaseCreateWalletInterface {
	return &usecaseCreateWallet{
		WalletRepo: walletRepo,
	}
}

func (u *usecaseCreateWallet) CreateWallet(in *usecasedto.InputCreateWalletDto) (out *usecasedto.OutputCreateWalletDto, err error) {
	if !u.WalletRepo.Create(in.Exchange, int64(in.Coin), in.Amount) {
		log.Println("UCW 00: ", err)
		return
	}

	wallet, err := u.WalletRepo.GetWithCoin(in.Exchange, in.Coin)
	if err != nil {
		log.Println("UCW 01: ", err)
		return
	}

	out = &usecasedto.OutputCreateWalletDto{
		Wallet:   wallet.Wallet,
		Exchange: wallet.Exchange,
		Coin:     wallet.Coin,
		Amount:   wallet.Amount,
	}

	return
}
