package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type tradeConfigRepository struct {
	Db database.DatabaseInterface
}

type TradeConfigRepositoryInterface interface {
	Create(*repositorydto.InputTradeConfigDto) bool
	List() (out []*repositorydto.OutputTradeConfigDto, err error)
}

func NewTradeConfigRepository(db database.DatabaseInterface) TradeConfigRepositoryInterface {
	return &tradeConfigRepository{
		Db: db,
	}
}

func (t *tradeConfigRepository) List() (out []*repositorydto.OutputTradeConfigDto, err error) {
	tx, err := t.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Panicln("TCRL 00: ", err)
	}

	stmt, err := tx.Prepare(`
		SELECT * FROM trade_config;
	`)

	if err != nil {
		log.Panicln("TCRL 01: ", err)
	}

	defer stmt.Close()

	res, err := stmt.Query()

	if err != nil {
		log.Panicln("UEKR 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var u repositorydto.OutputTradeConfigDto

		err = res.Scan(
			&u.TradeConfig,
			&u.User,
			&u.Modality,
			&u.Strategy,
			&u.StrategyVariant,
			&u.Parity,
			&u.Exchange,
			&u.OperationQuantity,
			&u.OperationAmount,
			&u.DefaultProfitPercentage,
			&u.WalletValueLimit,
			&u.Enabled,
		)

		if err != nil {
			log.Panicln("UEKR 03: ", err)
			return
		}

		out = append(out, &u)
	}

	return
}

func (t *tradeConfigRepository) Create(in *repositorydto.InputTradeConfigDto) bool {
	tx, err := t.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Panicln("TCRC 00: ", err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO trade_config (
			modality,
			user,
			strategy,
			strategy_variant,
			parity,
			exchange,
			operation_quantity,
			operation_amount,
			default_profit_percentage,
			wallet_value_limit,
			enabled
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`)

	if err != nil {
		log.Panicln("TCRC 01: ", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		in.Modality,
		in.User,
		in.Strategy,
		in.StrategyVariant,
		in.Parity,
		in.Exchange,
		in.OperationQuantity,
		in.OperationAmount,
		in.DefaultProfitPercentage,
		in.WalletValueLimit,
		in.Enabled,
	)

	if err != nil {
		log.Panicln("TCRC 02: ", err)
		return false
	}

	err = tx.Commit()
	if err != nil {
		log.Panicln("TCRC 03: ", err)
		return false
	}

	return true
}
