package model

type Config struct {
	BaseModel
	SteamPath      string `json:"steam_path"`
	DayzPath       string `json:"dayz_path"`
	DayzServerPath string `json:"dayz_server_path"`
	WorkshopPath   string `json:"workshop_path"`
}
