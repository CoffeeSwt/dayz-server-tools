package app

var bindedObjects []interface{}

// GetBind returns the binded objects for the app
func GetBind() []interface{} {
	var customBind []interface{}
	customBind = append(customBind, bindedObjects...)
	return customBind
}
