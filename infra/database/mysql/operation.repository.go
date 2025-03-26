package mysql

import (
	"context"
	"database/sql"
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
	Create(userId uint64, parity int64, exchange, strategy, strategyVariant int64, investedAmount float64, enabled bool) int64
	Count(userId uint64, parity int64) int64
	List(in *repositorydto.InputOperationDto) (list []*repositorydto.OjoinOMDFT)
	Get(operation int64) (out *repositorydto.OutputOperationDto)
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

func (o *operationRepository) Delete(operation uint64) bool {
	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("ORD 00: ", err)
		return false
	}
	query := `DELETE operation WHERE operation = ?`

	_, err = tx.Exec(query, operation)
	if err != nil {
		log.Println("ORD 01: ", err)
		return false
	}

	tx.Commit()
	return true
}

func (o *operationRepository) UpdateDynamically(updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool {
	_, wheres, updates, _ := utils.QueryFormatter(updateFields, updatefieldValues, wherecolumns, wherevalues, paginationValues, []string{}, []any{}, order)
	query := `UPDATE operation SET ` + updates + wheres

	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("ORUD 00: ", err)
		return false
	}
	_, err = tx.Exec(query)

	if err != nil {
		log.Println("ORUD 01: ", err)
		return false
	}

	tx.Commit()
	return true
}

func (o *operationRepository) Update(p *repositorydto.InputOperationDto) bool {
	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("ORU 00: ", err)
		return false
	}

	stmt, err := tx.Prepare(`
		UPDATE operation 
		SET mts_start = ?, mts_finish = ?, open_price = ?, close_price = ?, invested_amount = ?, profit_amount = ?, profit = ?, closed = ?, audit = ?, enabled = ?
		WHERE operation = ?
	`)

	if err != nil {
		log.Panicln("ORU 01: ", err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.MtsStart, p.MtsFinish, p.OpenPrice, p.ClosePrice, p.InvestedAmount, p.ProfitAmount, p.Profit, p.Closed, p.Audit, p.Enabled, p.Operation)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("ORU 02: ", err)
		return false
	}

	tx.Commit()
	return true

}

func (o *operationRepository) ListByPeriod(from, to int64) (list []*repositorydto.OutputOperationWStrategyDto) {
	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("ORLBP 00: ", err)
		return
	}

	stmt, err := tx.Prepare(`
		SELECT o.operation, o.user, o.parity, o.exchange, o.strategy, s.name as strategy_name, o.strategy_variant, sv.name as strategy_variant_name, o.mts_start, o.mts_finish, o.profit, o.profit_amount, o.invested_amount, o.open_price, o.close_price, o.closed, o.audit, o.enabled
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

		err = res.Scan(&o.Operation, &o.User, &o.Parity, &o.Exchange, &o.Strategy, &o.StrategyName, &o.StrategyVariant, &o.StrategyVariantName, &o.MtsStart, &o.MtsFinish, &o.Profit, &o.ProfitAmount, &o.InvestedAmount, &o.OpenPrice, &o.ClosePrice, &o.Closed, &o.Audit, &o.Enabled)

		if err != nil {
			log.Panic("ORLBP 03: ", err)
			return
		}

		list = append(list, &o)
	}

	tx.Commit()
	return
}

func (o *operationRepository) ListAll() (list []*repositorydto.OutputOperationWStrategyDto) {
	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("ORLA 00: ", err)
		return
	}

	stmt, err := tx.Prepare(`
		SELECT o.operation, o.user, o.parity, o.exchange, o.strategy, s.name as strategy_name, o.strategy_variant, sv.name as strategy_variant_name, o.mts_start, o.mts_finish, o.profit, o.invested_amount, o.open_price, o.close_price, o.closed, o.audit, o.enabled
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

		err = res.Scan(&o.Operation, &o.User, &o.Parity, &o.Exchange, &o.Strategy, &o.StrategyName, &o.StrategyVariant, &o.StrategyVariantName, &o.MtsStart, &o.MtsFinish, &o.Profit, &o.InvestedAmount, &o.OpenPrice, &o.ClosePrice, &o.Closed, &o.Audit, &o.Enabled)

		if err != nil {
			log.Panic("ORLA 03: ", err)
			return
		}

		list = append(list, &o)
	}

	tx.Commit()
	return
}

func (o *operationRepository) List(in *repositorydto.InputOperationDto) (list []*repositorydto.OjoinOMDFT) {
	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("ORL 00: ", err)
		return
	}

	stmt, err := tx.Prepare(`
		SELECT o.operation, o.user, o.parity, o.exchange, o.strategy, o.strategy_variant, o.mts_start, o.mts_finish, o.profit, o.invested_amount, o.open_price, o.close_price, o.closed, o.audit, o.enabled, om.operation_meta_data_fast_trade, om.minimum_price, om.maximum_price, om.last_price, om.is_receding
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

		err = res.Scan(&o.Operation, &o.User, &o.Parity, &o.Exchange, &o.Strategy, &o.StrategyVariant, &o.MtsStart, &o.MtsFinish, &o.Profit, &o.InvestedAmount, &o.OpenPrice, &o.ClosePrice, &o.Closed, &o.Audit, &o.Enabled, &o.OperationMetaDataFastTrade, &o.MinimumPrice, &o.MaximumPrice, &o.LastPrice, &o.IsReceding)

		if err != nil {
			log.Panic("ORL 03: ", err)
			return
		}

		list = append(list, &o)
	}

	tx.Commit()
	return
}

func (o *operationRepository) Get(operation int64) (out *repositorydto.OutputOperationDto) {
	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("ORF 00: ", err)
		return
	}

	stmt, err := tx.Prepare(`
		SELECT o.operation, o.user, o.parity, o.exchange, o.strategy, o.strategy_variant, o.mts_start, o.mts_finish, o.profit, o.invested_amount, o.open_price, o.close_price, o.closed, o.audit, o.enabled
		FROM operation o
		WHERE o.operation = ?
	`)

	if err != nil {
		log.Panic("ORF 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(operation).Scan(&out.Operation, &out.User, &out.Parity, &out.Exchange, &out.Strategy, &out.StrategyVariant, &out.MtsStart, &out.MtsFinish, &out.Profit, &out.InvestedAmount, &out.OpenPrice, &out.ClosePrice, &out.Closed, &out.Audit, &out.Enabled)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("ORL 02: ", err)
		return
	}

	tx.Commit()
	return
}

func (o *operationRepository) Count(userId uint64, parity int64) (count int64) {
	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("ORCO 00: ", err)
		return
	}

	stmt, err := tx.Prepare(`SELECT COUNT(*) FROM operation WHERE closed = 0`)

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

	tx.Commit()
	return
}

func (o *operationRepository) Create(userId uint64, parity int64, exchange, strategy, strategyVariant int64, investedAmount float64, enabled bool) int64 {
	tx, err := o.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println("ORCO 00: ", err)
		return 0
	}

	stmt, err := tx.Prepare(`INSERT INTO operation(user, parity, exchange, strategy, strategy_variant, mts_start, invested_amount, enabled) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Panicln("ORC 01: ", err)
		return 0
	}

	defer stmt.Close()

	res, err := stmt.Exec(userId, parity, exchange, strategy, strategyVariant, time.Now().UnixMicro(), investedAmount, enabled)

	if err != nil {
		log.Panicln("ORC 02: ", err)
		return 0
	}

	li, _ := res.LastInsertId()

	tx.Commit()
	return li
}
