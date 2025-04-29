package di

import (
	"DataTask/internal/config"
	"DataTask/internal/repository/database"
	"DataTask/pkg/logger"
	"database/sql"
	"time"

	"fmt"
)

func InitializeDatabase(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Base,
	)

	logger.Log.Info("Try connect to pg:", dsn)

	var db *sql.DB
	var err error

	maxAttempts := 10
	for i := 1; i <= maxAttempts; i++ {
		db, err = database.ConnectPostgres(dsn)
		if err == nil && db.Ping() == nil {
			return db, nil
		}
		logger.Log.Warnf("DB not ready (attempt %d/%d), retrying in 2s...\n", i, maxAttempts)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("database not available after %d attempts: %w", maxAttempts, err)
}
