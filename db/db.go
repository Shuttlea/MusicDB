package db

import (
	"fmt"
	"log/slog"
	"music_db/config"
	sw "music_db/go"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s database=%s port=%s sslmode=%s password=%s", cfg.HostDB, cfg.UserDB, cfg.DataBase, cfg.PortDB, cfg.SslmodeDB, cfg.PasswordDB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Debug("Can't open database", "Error", err)
		panic(err)
	}
	err = db.AutoMigrate(&sw.Song{})
	if err != nil {
		slog.Debug("Can't open database", "Error", err)
		panic(err)
	}
	return db
}

func CloseDBConnection(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
