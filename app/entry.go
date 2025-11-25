package app

import "dayz-server-tools/db"

func GetBind() []interface{} {
	var customBind []interface{}
	customBind = append(customBind, db.GetOrm())
	return customBind
}
