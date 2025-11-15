package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runsystem-test/internal/entity"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func ConnectDB() (*gorm.DB, error) {
	LoadEnv()
	dsn := os.Getenv("DB_PATH")
	if dsn == "" {
		return nil, fmt.Errorf("DB_PATH environment is not set")
	}

	dir := filepath.Dir(dsn)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, fmt.Errorf("Failed to create directory: %w", err)
		}
	}

	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        dsn,
	}, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	if err := db.AutoMigrate(&entity.UserEntity{}); err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %w", err)
	}

	return db, nil
}
