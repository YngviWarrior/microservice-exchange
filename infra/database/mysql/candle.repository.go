package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
)

type candleRepository struct {
	Db database.DatabaseInterface
}

type CandleRepositoryInterface interface {
	GetLastPrice(parity, exchange uint64) (close string)
	GetLastPriceConn(parity, exchange uint64) (close string)
	ListLimit(parity, exchange, limit int64) (list []*repositorydto.OutputCandleDto)
	GetFirstMts(parity, exchange int64) (c repositorydto.OutputCandleDto)
	GetFirst(parity, exchange, from int64) (c repositorydto.OutputCandleDto)
	GetAvg(from, to int64) (list []*repositorydto.OutputCandleDto)
	Create(candle *repositorydto.InputCandleDto) bool
	CreateDirect(candle *repositorydto.InputCandleDto) bool
	CreateList(candles []*repositorydto.InputCandleDto) bool
	GetFirstPrice(parity, exchange, from uint64) (c repositorydto.OutputCandleDto)
}

func NewCandleRepository(db database.DatabaseInterface) CandleRepositoryInterface {
	return &candleRepository{
		Db: db,
	}
}

func (t *candleRepository) GetFirstPrice(parity, exchange, from uint64) (c repositorydto.OutputCandleDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT parity, exchange, close, mts
		FROM candle 
		WHERE mts > ? AND parity = ? AND exchange = ?
		LIMIT 1
	`)

	if err != nil {
		log.Panicln("CFF 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(from, parity, exchange).Scan(&c.Parity, &c.Exchange, &c.Close, &c.Mts)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CFF 02: ", err)
		return
	}

	return
}

func (t *candleRepository) GetLastPriceConn(parity, exchange uint64) (close string) {
	stmt, err := t.Db.GetDatabase().PrepareContext(context.TODO(), `
		SELECT close
		FROM candle 
		WHERE parity = ? AND exchange = ?
		ORDER BY mts DESC 
		LIMIT 1
	`)

	if err != nil {
		log.Panicln("CRFLPC 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(parity, exchange).Scan(&close)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRFLPC 02: ", err)
		return
	}

	return
}

func (t *candleRepository) GetLastPrice(parity, exchange uint64) (close string) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT close
		FROM candle 
		WHERE parity = ? AND exchange = ?
		ORDER BY mts DESC 
		LIMIT 1
	`)

	if err != nil {
		log.Panicln("CRFLP 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(parity, exchange).Scan(&close)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRFLP 02: ", err)
		return
	}

	return
}

func (t *candleRepository) CreateList(candles []*repositorydto.InputCandleDto) bool {
	if len(candles) == 0 {
		return true
	}

	query := `INSERT INTO candle(parity, exchange, mts, open, close, high, low, volume) VALUES `
	values := []string{}
	args := []any{}

	for _, candle := range candles {
		values = append(values, "(?, ?, ?, ?, ?, ?, ?, ?)")
		args = append(args, candle.Parity, candle.Exchange, candle.Mts*1000, candle.Open, candle.Close, candle.High, candle.Low, candle.Volume)
	}

	query += strings.Join(values, ",")
	query += ` ON DUPLICATE KEY UPDATE 
    parity = VALUES(parity), 
    exchange = VALUES(exchange), 
    mts = VALUES(mts), 
    open = VALUES(open), 
    close = VALUES(close), 
    high = VALUES(high), 
    low = VALUES(low), 
    volume = VALUES(volume)`

	stmt, err := t.Db.GetDatabase().Prepare(query)

	if err != nil {
		log.Panicln("CRC 01: ", err)
		return false
	}

	defer stmt.Close()
	_, err = stmt.Exec(args...)

	if err != nil {
		log.Panicln("CRC 02: ", err)
		return false
	}

	return true
}

func (t *candleRepository) Create(candle *repositorydto.InputCandleDto) bool {
	constraint := `
		ON DUPLICATE KEY UPDATE 
		parity = ` + fmt.Sprintf("%v", candle.Parity) + `,
		exchange = ` + fmt.Sprintf("%v", candle.Exchange) + `,
		open = ` + fmt.Sprintf("%v", candle.Open) + `,
		close = ` + fmt.Sprintf("%v", candle.Close) + `,
		high = ` + fmt.Sprintf("%v", candle.High) + `,
		low = ` + fmt.Sprintf("%v", candle.Low) + `,
		volume = ` + fmt.Sprintf("%v", candle.Volume) + `
	`
	query := `INSERT INTO candle(parity, exchange, mts, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?, ?, ?)` + constraint

	stmt, err := t.Db.GetDatabase().Prepare(query)

	if err != nil {
		log.Panicln("CRC 01: ", err)
		return false
	}

	defer stmt.Close()
	_, err = stmt.Exec(candle.Parity, candle.Exchange, (candle.Mts * 1000), candle.Open, candle.Close, candle.High, candle.Low, candle.Volume)

	if err != nil {
		log.Panicln("CRC 02: ", err)
		return false
	}

	return true
}

func (t *candleRepository) CreateDirect(candle *repositorydto.InputCandleDto) bool {
	constraint := `
		ON DUPLICATE KEY UPDATE 
		parity = ` + fmt.Sprintf("%v", candle.Parity) + `,
		exchange = ` + fmt.Sprintf("%v", candle.Exchange) + `,
		open = ` + fmt.Sprintf("%v", candle.Open) + `,
		close = ` + fmt.Sprintf("%v", candle.Close) + `,
		high = ` + fmt.Sprintf("%v", candle.High) + `,
		low = ` + fmt.Sprintf("%v", candle.Low) + `,
		volume = ` + fmt.Sprintf("%v", candle.Volume) + `
	`
	query := `INSERT INTO candle(parity, exchange, mts, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?, ?, ?)` + constraint

	stmt, err := t.Db.GetDatabase().Prepare(query)

	if err != nil {
		log.Panicln("CRCD 01: ", err)
		return false
	}

	defer stmt.Close()
	_, err = stmt.Exec(candle.Parity, candle.Exchange, (candle.Mts * 1000), candle.Open, candle.Close, candle.High, candle.Low, candle.Volume)

	if err != nil {
		log.Panicln("CRCD 02: ", err)
		return false
	}

	return true
}

func (t *candleRepository) ListLimit(parity, exchange, limit int64) (list []*repositorydto.OutputCandleDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT parity, exchange, close, mts
		FROM candle 
		WHERE parity = ? AND exchange = ?
		ORDER BY mts DESC
		LIMIT ?
	`)

	if err != nil {
		log.Panicln("CRLL 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(parity, exchange, limit)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRLL 02: ", err)
		return
	}

	for res.Next() {
		var c repositorydto.OutputCandleDto

		err := res.Scan(&c.Parity, &c.Exchange, &c.Close, &c.Mts)

		if err != nil {
			log.Panic("CRLL 03: ", err)
		}

		list = append(list, &c)
	}

	return
}

func (t *candleRepository) GetFirstMts(parity, exchange int64) (out repositorydto.OutputCandleDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT parity, exchange, close, mts
		FROM candle 
		WHERE parity = ? AND exchange = ?
		ORDER BY mts DESC
		LIMIT 1
	`)

	if err != nil {
		log.Panicln("CRFFM 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(parity, exchange).Scan(&out.Parity, &out.Exchange, &out.Close, &out.Mts)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRFFM 02: ", err)
		return
	}

	return
}

func (t *candleRepository) GetFirst(parity, exchange, from int64) (out repositorydto.OutputCandleDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT parity, exchange, close, mts
		FROM candle 
		WHERE mts > ? AND parity = ? AND exchange = ?
		LIMIT 1
	`)

	if err != nil {
		log.Panicln("CRFF 01: ", err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(from, parity, exchange).Scan(&out.Parity, &out.Exchange, &out.Close, &out.Mts)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRFF 02: ", err)
		return
	}

	return
}

func (t *candleRepository) GetAvg(from, to int64) (list []*repositorydto.OutputCandleDto) {
	stmt, err := t.Db.GetDatabase().Prepare(`
		SELECT parity, exchange, AVG(close), (((MIN(close) - MAX(close)) / MAX(close)) * 100) as roc
		FROM candle 
		WHERE mts BETWEEN ? AND ?
		GROUP BY parity, exchange
	`)

	if err != nil {
		log.Panicln("CRFA 01: ", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(from, to)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		log.Panicln("CRFA 02: ", err)
		return
	}

	defer res.Close()

	for res.Next() {
		var c repositorydto.OutputCandleDto

		err := res.Scan(&c.Parity, &c.Exchange, &c.Close, &c.Roc)

		if err != nil {
			log.Panic("CRFA 03: ", err)
		}

		list = append(list, &c)
	}

	return
}
