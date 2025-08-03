package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"user-service/internal/config"
)

type PostgresDB struct {
	dbPool *pgxpool.Pool
}

var (
	ErrPostgresConnection  = errors.New("error connecting to postgres database")
	ErrPostgresParseConfig = errors.New("error parsing postgres config")
	ErrPostgresPing        = errors.New("error pinging database")
)

func NewPostgresDB(ctx context.Context, cfg config.DBConfig) (*PostgresDB, error) {

	// Parse Option
	pgxCfg, err := pgxpool.ParseConfig(cfg.Uri)
	if err != nil {
		return nil, ErrPostgresParseConfig
	}

	// Setup options
	pgxCfg.MaxConns = int32(cfg.MaxConn)
	pgxCfg.MinConns = int32(cfg.MinConn)
	pgxCfg.MaxConnLifetime = cfg.MaxConnLifeTime
	pgxCfg.MaxConnIdleTime = cfg.MaxIdleTime
	pgxCfg.HealthCheckPeriod = cfg.HealthCheckPeriod

	// Create Connection
	dbPool, err := pgxpool.New(ctx, cfg.Uri)
	if err != nil {
		return nil, ErrPostgresConnection
	}

	// Checking the connection
	if err := dbPool.Ping(ctx); err != nil {
		return nil, ErrPostgresPing
	}

	return &PostgresDB{
		dbPool: dbPool,
	}, nil
}
