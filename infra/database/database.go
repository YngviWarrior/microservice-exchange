package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseInterface interface {
	CreateConnection() *sql.Conn
}

type database struct {
	Pool *sql.DB
}

func NewDatabase() DatabaseInterface {
	var db *sql.DB
	var err error

	switch os.Getenv("ENVIROMENT") {
	case "local", "testnet":
		db, err = sql.Open("mysql", os.Getenv("DB_LOCAL"))
	default:
		db, err = sql.Open("mysql", os.Getenv("DB_SERVER"))
	}

	if err != nil {
		log.Fatal("ND 01: ", err.Error())
	}

	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxOpenConns(0)

	return &database{
		Pool: db,
	}
}

func (d *database) CreateConnection() *sql.Conn {
	conn, err := d.Pool.Conn(context.Background())
	if err != nil {
		log.Fatal("CC 01: ", err.Error())
	}

	return conn
}

func (d *database) CreateTransaction() *sql.Tx {
	tx, err := d.Pool.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Fatal("CT 01: ", err.Error())
	}

	return tx
}
