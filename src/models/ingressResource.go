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
	Annotations map[string]string
}

// ingress特殊参数
type IngressAnnotation struct {
	Name    string `json:"name"`
	Remarks string `json:"remarks"`
	Description string `json:"description"`
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
	IngressAnnotationsHead = "nginx.ingress.kubernetes.io/ingress.class"
)

var IngressAnnotationsitem = []IngressAnnotation{
	IngressAnnotation{"ingress.class", "绑定控制器(*必须填项目)","（*必填项目）ingress-class值(参见internal/ingress/annotations/class/main.go中的IsValid方法)它将只处理一个未设置的类注释否则就需要类注释。"},
	IngressAnnotation{"app-root", "重定向","如果Application Root公开在不同的路径中需要重定向，那么将注释nginx.ingress.kubernetes.io/app-Root设置为重定向/的请求"},
	IngressAnnotation{"affinity", "会话亲和度","可以在入口的所有 Upstreams 中启用和设置亲和类型。通过这种方式，请求将总是定向到相同的上游服务器。唯一可用于 NGINX 的亲和类型是 cookie注意如果一个主机定义了多个入口，并且至少有一个入口使用了 nginx.Ingress.kubernetes.io/affinity : cookie，那么只有使用 nginx.Ingress.kubernetes.io/affinity 的入口路径才会使用会话 cookie 关联。通过随机选择后端服务器，所有在主机的其他接口上定义的路径都将得到负载平衡。"},
	IngressAnnotation{"affinity-mode", "会话的粘性","定义了会话的粘性。将其设置为均衡(默认值)将在部署扩展时重新分配某些会话，从而重新平衡服务器上的负载。将其设置为持久不会将会话重新平衡到新的服务器，因此提供了最大的粘性。"},
	IngressAnnotation{"auth-realm", "区域认证","认证请求的范围：例子auth-realm：Authentication Required - foo"},
	IngressAnnotation{"auth-secret", "认证K8s密钥","Secret 的名称，其中包含用户名和密码，这些用户名和密码被授予访问 Ingress 规则中定义的路径的权限。这个注释还接受“ namespace/secretName”的替代形式，在这种情况下，Secret 查找在被引用的名称空间而不是 Ingress 名称空间中执行。"},
	IngressAnnotation{"auth-secret-type", "认证密钥类型(auth-file|auth-map)","用来指定Secret的认证类型[auth-file|auth-map]"},
	IngressAnnotation{"auth-type", "HTTP身份验证类型(basic|digest)","指示 HTTP 身份验证类型: 基本或 HTTP摘要认证。"},
	IngressAnnotation{"auth-tls-secret", "包含完整证书Secret的名称","包含完整证书颁发机构链的 Secret 的名称ca.crt能够对这个入口进行验证。这个注释还接受“ namespace/secretName”的替代形式，在这种情况下，Secret 查找在被引用的名称空间而不是 Ingress 名称空间中执行"},
	IngressAnnotation{"auth-tls-verify-depth","认证的深度(Number)"," 提供的客户端证书与证书颁发机构链之间的验证深度"},
	IngressAnnotation{"auth-tls-error-page","证书认证错误时候出现的错误页","发生证书认证错误时，应重定向用户的 URL/页面,例子：auth-tls-error-page：http://www.mysite.com/error-cert.html"},
	IngressAnnotation{"auth-tls-verify-client","开启验证true/false","允许验证客户端证书"},
	IngressAnnotation{"auth-tls-pass-certificate-to-upstream","默认不启用"," 指示是否将接收到的证书传递给上游服务器。默认情况下禁用此选项"},
}

