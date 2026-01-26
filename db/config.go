package db

type Config struct {
	BaseModel
	ServerToolsAbsPath string `json:"server_tools_abs_path"` // 服务器工具绝对路径

	SteamPath      string `json:"steam_path"`       // steam路径
	DayZPath       string `json:"dayz_path"`        // dayz路径
	DayZServerPath string `json:"dayz_server_path"` // dayz服务器路径
}
