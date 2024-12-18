package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

type DbImpl struct {
	DB *sql.DB
}

func NewDbImpl(db *sql.DB) *DbImpl {
	return &DbImpl{
		DB: db,
	}
}

func (dbIMPL *DbImpl) StartTX(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := dbIMPL.DB.BeginTx(ctx, opts)
	return tx, err
}

func (dbIMPL *DbImpl) RunWithTransaction(ctx context.Context, opts *sql.TxOptions, fn func(*sql.Tx) error) error {
	tx, err := dbIMPL.StartTX(ctx, opts)
	if err != nil {
		log.Err(err).Msg("cannot start tx")
		return err
	}

	err = fn(tx)
	if err != nil {
		log.Err(err).Msg("roolback")
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func NewDB() *sql.DB {
	
const(
   host 	= "q4jei.h.filess.io"
   port     = "3306"
   user     = "db2025_satisfied"
   password = "6ae5f03b9845cf7eaaabfb35ccc63c9e934ce940"
   dbname   = "db2025_satisfied"
)

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err := sql.Open("mysql", uri)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

