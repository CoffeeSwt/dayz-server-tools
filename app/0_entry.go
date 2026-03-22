package app

// GetBind returns the binded objects for the app
func GetBind() []interface{} {

	return []interface{}{
		GetServer(),
	}
}
