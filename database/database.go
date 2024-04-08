package database

import (
	"context"
	"database/sql"
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
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/hey_notes_api?parseTime=true")
	
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

