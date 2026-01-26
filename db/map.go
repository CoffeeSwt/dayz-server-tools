package db

type Map struct {
	BaseModel
	Name        string `json:"name"`
	MapChinese  string `json:"map_chinese"`
	MissionName string `json:"mission_name"`
	Tips        string `json:"tips"`
	Image       string `json:"image"`
}

func GetMapInit() func() {
	maps := []Map{
		{
			Name:        "chernarusplus",
			MapChinese:  "切尔那鲁斯",
			MissionName: "dayzOffline.chernarusplus",
			Tips:        "官方默认地图",
			Image:       "/static/images/chernarusplus.jpg",
		},
		{
			Name:        "enoch",
			MapChinese:  "利沃利亚",
			MissionName: "dayzOffline.enoch",
			Tips:        "官方DLC地图",
			Image:       "/static/images/enoch.jpg",
		},
		{
			Name:        "sakhal",
			MapChinese:  "萨哈尔",
			MissionName: "dayzOffline.sakhal",
			Tips:        "官方DLC地图",
			Image:       "/static/images/sakhal.jpg",
		},
	}
	return func() {
		for _, m := range maps {
			go _db.FirstOrCreate(&Map{
				Name:        m.Name,
				MapChinese:  m.MapChinese,
				MissionName: m.MissionName,
				Tips:        m.Tips,
				Image:       m.Image,
			}, Map{
				Name: m.Name,
			})
		}
	}
}
