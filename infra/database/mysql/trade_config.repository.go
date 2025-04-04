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
	Update(in *repositorydto.InputTradeConfigDto) bool
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
		SELECT tc.trade_config, tc.user, tc.modality, tc.strategy, tc.strategy_variant, tc.parity, tc.exchange,
			tc.operation_quantity, tc.operation_amount, tc.enabled, tc.default_profit_percentage, tc.wallet_value_limit,
			m.name modality_name, s.name strategy_name, s.enabled strategy_enabled, sv.name strategy_variant_name, sv.enabled as strategy_variant_enabled, 
			p.symbol symbol_name, e.name exchange_name
		FROM trade_config tc
		JOIN modality m ON m.modality = tc.modality
		JOIN strategy s ON s.strategy = tc.strategy
		JOIN strategy_variant sv ON sv.strategy_variant = tc.strategy_variant
		JOIN parity p  ON p .parity = tc.parity
		JOIN exchange e ON e.exchange = tc.exchange
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
			&u.Enabled,
			&u.DefaultProfitPercentage,
			&u.WalletValueLimit,
			&u.ModalityName,
			&u.StrategyName,
			&u.StrategyEnabled,
			&u.StrategyVariantName,
			&u.StrategyVariantEnabled,
			&u.ParitySymbol,
			&u.ExchangeName,
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

func (o *tradeConfigRepository) Update(in *repositorydto.InputTradeConfigDto) bool {
	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("TRU 00: ", err)
		return false
	}

	stmt, err := tx.Prepare(`
		UPDATE trade_config 
		SET modality = ?,
			user = ?,
			strategy = ?,
			strategy_variant = ?,
			parity = ?,
			exchange = ?,
			operation_quantity = ?,
			operation_amount = ?,
			default_profit_percentage = ?,
			wallet_value_limit = ?,
			enabled = ?
		WHERE trade_config = ?
	`)

	if err != nil {
		log.Panicln("TRU 01: ", err)
		return false
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
		in.TradeConfig,
	)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("TRU 02: ", err)
		return false
	}

	tx.Commit()
	return true

}
