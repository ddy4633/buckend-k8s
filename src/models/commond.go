package models

// 数据返回的格式
type ReturnMsg struct {
	Type string      `json:"type"`
	Ns   string      `json:"ns"`
	Data interface{} `json:"data"`
}
