package db

type Server struct {
	BaseModel
	Name       string `json:"name" gorm:"unique"` //服务器名称
	Map        string `json:"map" gorm:"unique"`  //地图名称
	MapChinese string `json:"map_chinese"`        //地图中文名称
}
