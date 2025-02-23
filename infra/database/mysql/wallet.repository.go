package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type walletRepository struct {
	Db database.DatabaseInterface
}

type WalletRepositoryInterface interface {
	List(*repositorydto.InputWalletDto) []*repositorydto.OutputWalletDto
	UpdateAmount(values string) bool
	Create(exchange uint64, coin int64, amount float64) bool
	GetWithCoin(exchange, coin uint64) (out repositorydto.OutputWalletWithCoinDto, err error)
	ListWithCoin(exchange uint64) (list []*repositorydto.OutputWalletWithCoinDto)
}

func NewWalletRepository(db database.DatabaseInterface) WalletRepositoryInterface {
	return &walletRepository{
		Db: db,
	}
}

func (m *walletRepository) List(in *repositorydto.InputWalletDto) (out []*repositorydto.OutputWalletDto) {
	tx, err := m.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Panicln("WRL 00 :", err)
	}

	stmt, err := tx.Prepare(`
		SELECT wallet, exchange, coin, amount
		FROM wallet
	`)

	if err != nil {
		log.Panicln("WRL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Panicln("WRL 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var u repositorydto.OutputWalletDto

		err := res.Scan(&u.Wallet, &u.Exchange, &u.Coin, &u.Amount)

		if err != nil {
			log.Panicln("WRL 03: ", err)
			return
		}

		out = append(out, &u)
	}

	tx.Commit()

	return
}

func (w *walletRepository) Create(exchange uint64, coin int64, amount float64) bool {
	tx, err := w.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("WRC 00: ", err)
		return false
	}

	query := `INSERT INTO wallet(exchange,  coin, amount) VALUES (?, ?, ?)`

	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println("WRC 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(exchange, coin, amount)

	if err != nil {
		log.Println("WRC 02: ", err)
		return false
	}

	tx.Commit()
	return true
}

func (w *walletRepository) GetWithCoin(exchange, coin uint64) (out repositorydto.OutputWalletWithCoinDto, err error) {
	tx, err := w.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("WRGWC 00: ", err)
		return
	}

	stmt, err := tx.Prepare(`
		SELECT w.wallet, w.exchange, w.coin, w.amount, c.symbol, c.active
		FROM wallet w
		JOIN coin c ON w.coin = c.coin AND c.active = 1
		WHERE w.exchange = ? AND w.coin = ?`)

	if err != nil {
		log.Panicln("WRGWC 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(exchange, coin).Scan(&out.Wallet, &out.Exchange, &out.Coin, &out.Amount, &out.Symbol, &out.Active)

	if err != nil && err != sql.ErrNoRows {
		log.Panicln("WRGWC 02: ", err)
		return
	}

	tx.Commit()
	return
}

func (w *walletRepository) ListWithCoin(exchange uint64) (list []*repositorydto.OutputWalletWithCoinDto) {
	tx, err := w.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("WRLWC 00: ", err)
		return
	}

	stmt, err := tx.Prepare(`
		SELECT w.wallet, w.exchange,  w.coin, w.amount, c.symbol, c.active
		FROM wallet w
		JOIN coin c ON w.coin = c.coin AND c.active = 1
		WHERE w.exchange = ?`)

	if err != nil {
		log.Panicln("WRLWC 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(exchange)

	if err != nil {
		log.Panicln("WRLWC 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var w repositorydto.OutputWalletWithCoinDto

		err := res.Scan(&w.Wallet, &w.Exchange, &w.Coin, &w.Amount, &w.Symbol, &w.Active)

		if err != nil {
			log.Panicln("WRLWC 03: ", err)
			return
		}

		list = append(list, &w)
	}

	tx.Commit()
	return
}

func (w *walletRepository) UpdateAmount(values string) bool {
	tx, err := w.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("WRUA 00: ", err)
		return false
	}

	query := `INSERT INTO wallet(wallet, exchange, coin, amount) VALUES ` + values
	query += ` ON DUPLICATE KEY UPDATE amount = VALUES(amount)`

	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println("WRUA 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		log.Println("WRUA 02: ", err)
		return false
	}

	tx.Commit()
	return true
}
