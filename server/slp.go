package server

type ServerLaunchParameters struct {
	Port       int    // 服务器端口
	Mission    string // 服务器任务文件
	Profiles   string // 服务器配置文件目录
	ClientMods string // 客户端mod目录
	ServerMods string // 服务器端mod目录
	Config     string // 服务器配置文件
}

func GetServerLaunchParameters() ServerLaunchParameters {
	return ServerLaunchParameters{
		Port:       2302,
		Mission:    "C:\\Server\\mpmissions\\empty.Bitterroot",
		Profiles:   "C:\\Server\\profiles\\Bitterroot",
		ClientMods: "",
		ServerMods: "",
		Config:     "C:\\Server\\serverCfgs\\Bitterroot.cfg",
	}
}
