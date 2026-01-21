package app

import "dayz-server-tools/db"

type DB struct {
}

func NewDB() *DB {
	return &DB{}
}

func (d *DB) TestDB() bool {
	db := db.GetDB()
	if db == nil {
		return false
	}
	return true
}
