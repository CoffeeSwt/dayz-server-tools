package db

import (
	"database/sql"
	"dayz-server-tools/config"
	"dayz-server-tools/logger"
	"os"
	"path/filepath"
	"strconv"
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

func GetSQLDB() *sql.DB {
	dbOnce.Do(initDB)
	s, _ := _db.DB()
	return s
}

func ensureDBFile(path string) error {
	if path == "" {
		path = config.GetDBPath()
	}
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

func initDB() {
	p := config.GetDBPath()
	if err := ensureDBFile(p); err != nil {
		panic(err)
	}
	lg := gormLogger.Default.LogMode(gormLogger.Info)
	if !config.IsDev() {
		lg = gormLogger.Default.LogMode(gormLogger.Warn)
	}
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
}

func applyPool(s *sql.DB) {
	maxOpen := parseIntEnv("DB_MAX_OPEN", 25)
	maxIdle := parseIntEnv("DB_MAX_IDLE", 25)
	life := parseDurationEnv("DB_CONN_MAX_LIFETIME", time.Hour)
	idleLife := parseDurationEnv("DB_CONN_MAX_IDLE_TIME", 30*time.Minute)
	s.SetMaxOpenConns(maxOpen)
	s.SetMaxIdleConns(maxIdle)
	s.SetConnMaxLifetime(life)
	s.SetConnMaxIdleTime(idleLife)
}

func parseIntEnv(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	if n, err := strconv.Atoi(v); err == nil {
		return n
	}
	return def
}

func parseDurationEnv(key string, def time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	if d, err := time.ParseDuration(v); err == nil {
		return d
	}
	return def
}
