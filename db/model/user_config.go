package model

type UserConfig struct {
	BaseModel
	Config Config `json:"config"`
}
