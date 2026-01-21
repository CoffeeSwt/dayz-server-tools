package db

import (
	"database/sql"
	"dayz-server-tools/logger"
	"dayz-server-tools/model"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	_db    *gorm.DB
	dbOnce sync.Once
)

func GetDB() *gorm.DB {
	dbOnce.Do(initDB)
	return _db
}

func initDB() {
	p := "database.db"
	if err := ensureDBFile(p); err != nil {
		panic(err)
	}
	lg := gormLogger.Default.LogMode(gormLogger.Info)
	gdb, err := gorm.Open(sqlite.Open(p), &gorm.Config{Logger: lg, PrepareStmt: false, SkipDefaultTransaction: true})
	if err != nil {
		logger.Error("DB init failed", "err", err)
		panic("failed to connect database")
	}
	_db = gdb
	sqlDB, err := _db.DB()
	if err != nil {
		panic(err)
	}
	applyPool(sqlDB)
	_db.Exec("PRAGMA foreign_keys = ON;")
	_db.Exec("PRAGMA journal_mode = WAL;")
	_db.Exec("PRAGMA busy_timeout = 5000;")

	s, err := _db.DB()
	if err != nil {
		panic(err)
	}
	if err := s.Ping(); err != nil {
		panic(err)
	}

	// 注册模型
	if err := model.RegisterModel(_db); err != nil {
		panic(err)
	}

}

func ensureDBFile(path string) error {
	dir := filepath.Dir(path)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		if cerr := f.Close(); cerr != nil {
			return cerr
		}
	}
	return nil
}

func applyPool(s *sql.DB) {
	s.SetMaxOpenConns(25)
	s.SetMaxIdleConns(25)
	s.SetConnMaxLifetime(1 * time.Hour)
	s.SetConnMaxIdleTime(30 * time.Minute)
}
