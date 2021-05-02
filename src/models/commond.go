package models

// 数据返回的格式
type ReturnMsg struct {
	Type string      `json:"Type"`
	Ns   string      `json:"Ns"`
	Data interface{} `json:"data"`
}
