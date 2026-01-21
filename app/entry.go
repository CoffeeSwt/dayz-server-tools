package app

import "dayz-server-tools/db"

var bindedObjects = []interface{}{}

func AppendModelObjects() []interface{} {
	bindedObjects = append(bindedObjects, db.GetModelList()...)
	return bindedObjects
}

// GetBind returns the binded objects for the app
func GetBind() []interface{} {
	var customBind []interface{}
	customBind = append(customBind, bindedObjects...)
	return customBind
}
