package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StorageConfig struct {
	username, password, host, port, database string
}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

func NewClient(ctx context.Context, maxAttems int, sc StorageConfig) (pool *pgxpool.Pool, err error) {

	// нужно сформировать строку подключения
	dsn := fmt.Sprintf("postgreqsl://%s:%s@%s:%s/%s", sc.username, sc.password, sc.host, sc.port, sc.database)
	err = DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttems, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgreSQL")
	}

	return pool, nil
}

func DoWithTries(fn func() error, attemps int, delay time.Duration) (err error) {

	for attemps < 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemps--

			continue
		}

		return nil
	}
	return
}
