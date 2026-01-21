package app

import "dayz-server-tools/db"

type DB struct {
}

func NewDB() *DB {
	AppendModelObjects()
	return &DB{}
}

func (d *DB) CheckConnection() bool {
	db := db.GetDB()
	_db, err := db.DB()
	if err != nil {
		return false
	}
	if err := _db.Ping(); err != nil {
		return false
	}
	return true
}
