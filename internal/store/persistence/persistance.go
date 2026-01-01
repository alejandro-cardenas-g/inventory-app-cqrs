package persistence

import (
	"context"
	"fmt"
	"inventory_cqrs/internal/config"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Persistance struct {
	TxManager *TxManager
	Products *ProductRepository
}

func NewStore(cfg config.StoreConfig) (*Persistance) {

	poolCfg, err := pgxpool.ParseConfig(cfg.URL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database initalized")

	poolCfg.MaxConns = int32(cfg.MaxOpenConns)
	poolCfg.MinConns = int32(cfg.MaxIdleConns)
	poolCfg.MaxConnLifetime = time.Duration(cfg.ConnMaxLifetime)
	poolCfg.MaxConnIdleTime = 1 * time.Minute
	poolCfg.HealthCheckPeriod = 30 * time.Second

	pool, err :=  pgxpool.NewWithConfig(context.Background(), poolCfg)

	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	store := &Persistance{}

	store.Products = NewProductRepository(pool)

	return store
}

type TxManager struct {
	pool *pgxpool.Pool
}

func NewTxManager(pool *pgxpool.Pool) *TxManager {
	return &TxManager{pool: pool}
}

func (tm *TxManager) WithTx(
	ctx context.Context,
	fn func(tx DBTX) error,
) error {
	tx, err := tm.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}