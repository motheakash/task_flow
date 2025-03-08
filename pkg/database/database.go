package database

import (
	"task_flow/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := "host=" + cfg.Host + " user=" + cfg.User + " password=" + cfg.Password + " dbname=" + cfg.DBName + " port=" + cfg.Port + " sslmode=" + cfg.SSLMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
