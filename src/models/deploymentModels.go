package models

// deployment模板
type Deployment struct {
	Name       string
	NameSpace  string
	Replicas   [3]int32 // 总副本数 可用副本数量 不可用副本数量
	Images     string
	IsComplete bool
	Message    string
	CreateTime string
	Pods       []*Pod
	Age        string
}

// namespace模板
type NsModel struct {
	Name string
}

// Pod的模板
type Pod struct {
	Name         string
	NameSpace    string
	Images       string
	NodeName     string   // 所在机器
	IP           []string // 当前POD IP
	Phase        string   // Pod当前所处阶段
	IsRead       bool     // 是否是就绪状态
	Message      string
	CreateTime   string
	RestartCount int32  // 重启次数
	Age          string // 创建时间
}

// Service模板
type Services struct {
	Name          string
	NameSpace     string
	NetworkType   interface{}
	Address       string
	ExportPort    int32
	ContainerPort interface{}
	Protocol      interface{} //通讯的方式
	Age           string
	CreateTime    string
}
