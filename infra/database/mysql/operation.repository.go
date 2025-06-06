package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	utils "github.com/YngviWarrior/microservice-exchange/utils"
)

type operationRepository struct {
	Db database.DatabaseInterface
}

type OperationRepositoryInterface interface {
	Create(in *repositorydto.InputOperationDto) int64
	Count(userId uint64, parity int64) int64
	List(in *repositorydto.InputOperationDto) (list []*repositorydto.OjoinOMDFT)
	ListEnabled() (list []*repositorydto.OjoinOMDFT)
	Get(operation uint64) (out repositorydto.OjoinOMDFT)
	ListAll() (list []*repositorydto.OutputOperationWStrategyDto)
	ListByPeriod(from, to int64) (list []*repositorydto.OutputOperationWStrategyDto)
	Update(p *repositorydto.InputOperationDto) bool
	UpdateDynamically(updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool
	Delete(operation uint64) bool
}

func NewOperationRepository(db database.DatabaseInterface) OperationRepositoryInterface {
	return &operationRepository{
		Db: db,
	}
}

func (t *operationRepository) Delete(operation uint64) bool {
	query := `DELETE operation WHERE operation = ?`

	_, err := t.Db.GetDatabase().Exec(query, operation)
	if err != nil {
		log.Println("ORD 01: ", err)
		return false
	}

	return true
}

func (t *operationRepository) UpdateDynamically(updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool {
	_, wheres, updates, _ := utils.QueryFormatter(updateFields, updatefieldValues, wherecolumns, wherevalues, paginationValues, []string{}, []any{}, order)
	query := `UPDATE operation SET ` + updates + wheres

	_, err := t.Db.GetDatabase().Exec(query)

	if err != nil {
		log.Println("ORUD 01: ", err)
		return false
	}

	return true
}

func (t *operationRepository) Update(p *repositorydto.InputOperationDto) bool {
	stmt, err := t.Db.GetDatabase().Prepare(`
		UPDATE operation 
		SET mts_start = ?, mts_finish = ?, open_price = ?, close_price = ?, invested_amount = ?, profit_amount = ?, profit = ?, closed = ?, audit = ?, enabled = ?, in_transaction = ?
		WHERE operation = ?
	`)

	if err != nil {
		log.Panicln("ORU 01: ", err)
		return false
	}
	fmt.Println("Update operation: ", p)
	defer stmt.Close()

	_, err = stmt.Exec(p.MtsStart, p.MtsFinish, p.OpenPrice, p.ClosePrice, p.InvestedAmount, p.ProfitAmount, p.Profit, p.Closed, p.Audit, p.Enabled, p.InTransaction, p.Operation)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("ORU 02: ", err)
		return false
	}

	return true

}

func (t *operationRepository) ListByPeriod(from, to int64) (list []*repositorydto.OutputOperationWStrategyDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT o.operation, o.user, o.parity, o.exchange, o.strategy, s.name as strategy_name, o.strategy_variant, sv.name as strategy_variant_name, o.mts_start, o.mts_finish, o.profit, o.profit_amount, o.invested_amount, o.open_price, o.close_price, o.closed, o.audit, o.enabled, o.in_transaction
		FROM operation o
		JOIN strategy s on o.strategy = s.strategy
		JOIN strategy_variant sv on o.strategy_variant = sv.strategy_variant
		WHERE o.mts_finish BETWEEN ? AND ?
	`)

	if err != nil {
		log.Panic("ORLBP 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(from, to)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("ORLBP 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var o repositorydto.OutputOperationWStrategyDto

		err = res.Scan(&o.Operation, &o.User, &o.Parity, &o.Exchange, &o.Strategy, &o.StrategyName, &o.StrategyVariant, &o.StrategyVariantName, &o.MtsStart, &o.MtsFinish, &o.Profit, &o.ProfitAmount, &o.InvestedAmount, &o.OpenPrice, &o.ClosePrice, &o.Closed, &o.Audit, &o.Enabled, &o.InTransaction)

		if err != nil {
			log.Panic("ORLBP 03: ", err)
			return
		}

		list = append(list, &o)
	}

	return
}

func (t *operationRepository) ListAll() (list []*repositorydto.OutputOperationWStrategyDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT o.operation, o.user, o.parity, o.exchange, o.strategy, s.name as strategy_name, o.strategy_variant, sv.name as strategy_variant_name, o.mts_start, o.mts_finish, o.profit, o.invested_amount, o.open_price, o.close_price, o.closed, o.audit, o.enabled, o.in_transaction
		FROM operation o
		JOIN strategy s ON o.strategy = s.strategy
		JOIN strategy_variant sv ON o.strategy_variant = sv.strategy_variant
	`)

	if err != nil {
		log.Panic("ORLA 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("ORLA 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var o repositorydto.OutputOperationWStrategyDto

		err = res.Scan(&o.Operation, &o.User, &o.Parity, &o.Exchange, &o.Strategy, &o.StrategyName, &o.StrategyVariant, &o.StrategyVariantName, &o.MtsStart, &o.MtsFinish, &o.Profit, &o.InvestedAmount, &o.OpenPrice, &o.ClosePrice, &o.Closed, &o.Audit, &o.Enabled, &o.InTransaction)

		if err != nil {
			log.Panic("ORLA 03: ", err)
			return
		}

		list = append(list, &o)
	}

	return
}

func (t *operationRepository) ListEnabled() (list []*repositorydto.OjoinOMDFT) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT o.operation, o.user, o.parity, p.symbol as parity_symbol, o.exchange, e.name as exchange_name, o.strategy, s.name as strategy_name, o.strategy_variant, sv.name as strategy_variant_name, o.mts_start, o.mts_finish, o.profit, o.invested_amount, o.open_price, o.close_price, o.closed, o.audit, o.enabled, o.in_transaction, om.operation_meta_data_fast_trade, om.minimum_price, om.maximum_price, om.last_price, om.is_receding
		FROM operation o
		INNER JOIN strategy s ON o.strategy = s.strategy
		INNER JOIN strategy_variant sv ON o.strategy_variant = sv.strategy_variant
		INNER JOIN parity p ON o.parity = p.parity
		INNER JOIN exchange e ON o.exchange = e.exchange
		LEFT JOIN operation_meta_data_fast_trade om ON o.operation = om.operation
		WHERE o.enabled = 1
	`)

	if err != nil {
		log.Panic("ORLE 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query()

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("ORLE 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var o repositorydto.OjoinOMDFT

		err = res.Scan(&o.Operation, &o.User, &o.Parity, &o.ParitySymbol, &o.Exchange, &o.ExchangeName, &o.Strategy, &o.StrategyName, &o.StrategyVariant, &o.StrategyVariantName, &o.MtsStart, &o.MtsFinish, &o.Profit, &o.InvestedAmount, &o.OpenPrice, &o.ClosePrice, &o.Closed, &o.Audit, &o.Enabled, &o.InTransaction, &o.OperationMetaDataFastTrade, &o.MinimumPrice, &o.MaximumPrice, &o.LastPrice, &o.IsReceding)

		if err != nil {
			log.Panic("ORLE 03: ", err)
			return
		}

		list = append(list, &o)
	}

	return
}

func (t *operationRepository) List(in *repositorydto.InputOperationDto) (list []*repositorydto.OjoinOMDFT) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT o.operation, o.user, o.parity, o.exchange, o.strategy, o.strategy_variant, o.mts_start, o.mts_finish, o.profit, o.invested_amount, o.open_price, o.close_price, o.closed, o.audit, o.enabled, o.in_transaction, om.operation_meta_data_fast_trade, om.minimum_price, om.maximum_price, om.last_price, om.is_receding
		FROM operation o
		LEFT JOIN operation_meta_data_fast_trade om ON o.operation = om.operation
		WHERE user = ? AND parity = ? AND exchange = ? 
		AND strategy = ? AND strategy_variant = ? 
		AND closed = ? AND enabled = ?
		
	`)

	if err != nil {
		log.Panic("ORL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(in.User, in.Parity, in.Exchange, in.Strategy, in.StrategyVariant, in.Closed, in.Enabled)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("ORL 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var o repositorydto.OjoinOMDFT

		err = res.Scan(&o.Operation, &o.User, &o.Parity, &o.Exchange, &o.Strategy, &o.StrategyVariant, &o.MtsStart, &o.MtsFinish, &o.Profit, &o.InvestedAmount, &o.OpenPrice, &o.ClosePrice, &o.Closed, &o.Audit, &o.Enabled, &o.InTransaction, &o.OperationMetaDataFastTrade, &o.MinimumPrice, &o.MaximumPrice, &o.LastPrice, &o.IsReceding)

		if err != nil {
			log.Panic("ORL 03: ", err)
			return
		}

		list = append(list, &o)
	}

	return
}

func (t *operationRepository) Get(operation uint64) (out repositorydto.OjoinOMDFT) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT o.operation, o.user, o.parity, o.exchange, o.strategy, o.strategy_variant, o.mts_start, o.mts_finish, o.profit, o.invested_amount, o.open_price, o.close_price, o.closed, o.audit, o.enabled, o.in_transaction
		FROM operation o
		WHERE o.operation = ?
	`)

	if err != nil {
		log.Panic("ORF 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(operation).Scan(&out.Operation, &out.User, &out.Parity, &out.Exchange, &out.Strategy, &out.StrategyVariant, &out.MtsStart, &out.MtsFinish, &out.Profit, &out.InvestedAmount, &out.OpenPrice, &out.ClosePrice, &out.Closed, &out.Audit, &out.Enabled, &out.InTransaction)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("ORL 02: ", err)
		return
	}

	return
}

func (t *operationRepository) Count(userId uint64, parity int64) (count int64) {
	stmt, err := t.Db.GetDatabase().Prepare(`SELECT COUNT(*) FROM operation WHERE closed = 0`)

	if err != nil {
		log.Panicln("ORCO 01: ", err)
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow().Scan(&count)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("ORCO 01: ", err)
		return
	}

	return
}

func (t *operationRepository) Create(in *repositorydto.InputOperationDto) int64 {
	stmt, err := t.Db.GetDatabase().Prepare(`INSERT INTO operation(user, parity, exchange, strategy, strategy_variant, mts_start, invested_amount, enabled, in_transaction) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Panicln("ORC 01: ", err)
		return 0
	}

	defer stmt.Close()

	res, err := stmt.Exec(in.User, in.Parity, in.Exchange, in.Strategy, in.StrategyVariant, time.Now().UnixMicro(), in.InvestedAmount, in.Enabled, in.InTransaction)

	if err != nil {
		log.Panicln("ORC 02: ", err)
		return 0
	}

	li, _ := res.LastInsertId()

	return li
}
