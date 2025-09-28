package sqlite

import (
	"context"
	"path/filepath"
	"poseidon/internal/common/path"
	"time"

	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	_dbFileName      = "poseidon.db"
	_maxOpenConns    = 1
	_maxIdleConns    = 1
	_connMaxLifetime = 5 * time.Minute
	_contextTimeout  = 2 * time.Second
)

func NewSQLite(lc fx.Lifecycle) (*gorm.DB, error) {
	dsn := filepath.Join(path.ProjectRootDir(), _dbFileName)
	gdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(_maxOpenConns)
	sqlDB.SetMaxIdleConns(_maxIdleConns)
	sqlDB.SetConnMaxLifetime(_connMaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), _contextTimeout)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return sqlDB.Close()
		},
	})
	return gdb, nil
}
