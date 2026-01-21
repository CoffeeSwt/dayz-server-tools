package model

type Config struct {
	BaseModel
	Theme string `json:"theme"` // 主题 'light' or 'dark'
}
