package models

// ingress对象
type Ingresses struct {
	Name       string
	NameSpace  string
	CreateTime string
	Labels     interface{}
	Status     string
	Rules      interface{}
	Address    interface{}
}

// ingress特殊参数
type Annotation struct {
	Name    string `json:"name"`
	Remarks string `json:"remarks"`
}

// create-Ingress 提交回来的Path信息
type IngressPath struct {
	Path    string `json:"path"`
	SvcName string `json:"svc_name"`
	Port    string `json:"port"`
}

// create-Ingress 提交回来的Rules信息
type IngressRules struct {
	Host  string         `json:"host"`
	Paths []*IngressPath `json:"paths"`
}

// annotations内部对象数据
type Rules struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 前端接收的annotation对象
type GetAnntation struct {
	Rules []*Rules `json:"rules"`
}

// 最终合并的对象
type IngressCreate struct {
	Name            string
	NameSpace       string
	AnnotationsData []*GetAnntation `json:"annotationsdata"`
	Rules           []*IngressRules
}

// 关于InGress Annotations的配置信息
const (
	IngressAnnotationsHead = "nginx.ingress.kubernetes.io/"
)

var IngressAnnotationsitem = []Annotation{
	Annotation{"ingress.class", "绑定控制器"},
	Annotation{"app-root", "重定向"},
	Annotation{"affinity", "会话亲和度"},
	Annotation{"affinity-mode", "会话的粘性"},
	Annotation{"auth-realm", "区域认证"},
	Annotation{"auth-secret", "认证K8s密钥"},
	Annotation{"auth-secret-type", "认证密钥类型(auth-file|auth-map)"},
	Annotation{"auth-type", "HTTP身份验证类型(basic|digest)"},
	Annotation{"auth-tls-secret", "包含完整证书Secret的名称"},
}

