package database

import (
	"fmt"
	"net/url"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
	sqlTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"

	"golang-core/api/src/common/parser/model"
)

type Db struct {
	logger  *zap.SugaredLogger
	primary *bun.DB
	replica *bun.DB
}

// Create new database context
func NewDbContext(cfg *model.DatabaseConfig, logger *zap.SugaredLogger) (db *Db, err error) {
	primary, err := dbConn(cfg, true)
	if err != nil {
		return db, err
	} else {
		logger.Info("connected to primary db ...")
	}
	replica, err := dbConn(cfg, false)
	if err != nil {
		return db, err
	} else {
		logger.Info("connected to replica db ...")
	}
	return &Db{
		logger:  logger,
		primary: primary,
		replica: replica,
	}, nil
}

// Create new database connection
func dbConn(cfg *model.DatabaseConfig, isPrimary bool) (*bun.DB, error) {
	var connectionString string
	if isPrimary {
		connectionString = fmt.Sprintf(
			"%s://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.Protocol,
			cfg.Username,
			url.QueryEscape(cfg.Password),
			cfg.URL,
			cfg.Port,
			cfg.Name,
		)
	} else {
		connectionString = fmt.Sprintf(
			"%s://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.Protocol,
			cfg.Username,
			url.QueryEscape(cfg.Password),
			cfg.ReplicaURL,
			cfg.Port,
			cfg.Name,
		)
	}

	pgConnector := pgdriver.NewConnector(
		pgdriver.WithDSN(connectionString),
		pgdriver.WithTimeout(30*time.Second),
	)
	sqlTrace.Register("pgdriver", pgConnector.Driver())
	conn := sqlTrace.OpenDB(pgConnector)
	conn.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime) * time.Second)
	conn.SetMaxOpenConns(cfg.MaxDBConns)
	conn.SetMaxIdleConns(cfg.MaxDBConns)
	conn.SetConnMaxIdleTime(time.Duration(cfg.MaxConnIdleTime) * time.Second)
	db := bun.NewDB(conn, pgdialect.New(), bun.WithDiscardUnknownColumns())
	err := db.Ping()
	return db, err
}

// Close the database context including primary and replica databases
func (db *Db) Close() error {
	err := db.primary.Close()
	if err != nil {
		db.logger.Errorf("failed to close primary DB: %w", err)
	}
	err = db.replica.Close()
	if err != nil {
		db.logger.Errorf("failed to close replica DB: %w", err)
	}
	return err
}

// Return primary db
func (db *Db) Primary() *bun.DB {
	return db.primary
}

// Return replica db
func (db *Db) Replica() *bun.DB {
	return db.replica
}
