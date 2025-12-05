package model

type Server struct {
	BaseModel
	Name       string `json:"name"`
	Map        string `json:"map"`
	MapChinese string `json:"map_chinese"`
}
