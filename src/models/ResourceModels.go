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
	Lables     map[string]string
	Pods       []*Pod
	Age        string
}

// namespace模板
type NsModel struct {
	Name string
}

// Pod的模板
type Pod struct {
	Name           string
	NameSpace      string
	Images         string
	NodeName       string   // 所在机器
	IP             []string // 当前POD IP
	Phase          string   // Pod当前所处阶段
	IsRead         bool     // 是否是就绪状态
	Message        string
	CreateTime     string
	Labels         map[string]string //标签
	ContainerPort  interface{}       // 内部使用端口
	Mount          interface{}
	RestartCount   int32  // 重启次数
	Age            string // 创建时间
	ContainersName []*Containers
	Annotation     map[string]string // 注解
	Tolerations    interface{}       // 容忍
	Secret         string            // 使用的权限
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

// Containers模板
type Containers struct {
	Name string
}

// Secret模板
type Secrets struct {
	Name       string
	NameSpace  string
	Type       string
	CreateTime string
}

var Secret_Type_Select = map[string]string{
	"Opaque":"自定义类型",
	"kubernetes.io/service-account-token":"服务账号令牌",
	"kubernetes.io/dockercfg":"docker配置",
	"kubernetes.io/dockerconfigjson":"docker配置(JSON)",
	"kubernetes.io/basic-auth":"Basic认证凭据",
	"kubernetes.io/ssh-auth":" SSH凭据",
	"kubernetes.io/tls":"TLS凭据",
	"bootstrap.kubernetes.io/token":"启动引导令牌数据",
}