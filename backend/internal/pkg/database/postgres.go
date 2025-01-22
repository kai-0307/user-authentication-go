package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunMigrations(dbURL string) error {
    m, err := migrate.New(
        "file://db/migrations",
        dbURL,
    )
    if err != nil {
        return fmt.Errorf("failed to create migrate instance: %v", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("failed to run migrations: %v", err)
    }

    return nil
}

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// SQLDBインスタンスを取得
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// コネクションプールの設定
	sqlDB.SetMaxIdleConns(10)  // アイドル状態のコネクションプール数
	sqlDB.SetMaxOpenConns(100) // オープンできる最大コネクション数

	return db, nil
}
