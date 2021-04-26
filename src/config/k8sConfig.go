package config

import (
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	"k8s-web/src/servers"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sConfig struct {
	DeployMentHandles *servers.DeployHandler    `inject:"-"`
	PodHandles        *servers.PodHandler       `inject:"-"`
	NamespaceHandles  *servers.NameSpaceHandler `inject:"-"`
	EventHandlers     *servers.EventHandler     `inject:"-"`
	ServiceHandlers   *servers.ServiceHandler   `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

// 初始化客户端
func (*K8sConfig) InitClient() *kubernetes.Clientset {
	rest, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		log.Println(err)
	}
	client, err := kubernetes.NewForConfig(rest)
	if err != nil {
		log.Println(err)
	}
	return client
}

// 初始化informer
func (this *K8sConfig) InitInformer() informers.SharedInformerFactory {
	fact := informers.NewSharedInformerFactory(this.InitClient(), 3*time.Second)

	// 初始化deployment的informer
	deployInformer := fact.Apps().V1().Deployments()
	deployInformer.Informer().AddEventHandler(this.DeployMentHandles)

	// 初始化Pod的informer
	podInformer := fact.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(this.PodHandles)

	// 初始化NameSpace
	nsInformer := fact.Core().V1().Namespaces()
	nsInformer.Informer().AddEventHandler(this.NamespaceHandles)

	// 初始化Service
	serverInformer := fact.Core().V1().Services()
	serverInformer.Informer().AddEventHandler(this.ServiceHandlers)

	// 初始化事件的信息
	EventInformer := fact.Core().V1().Events()
	EventInformer.Informer().AddEventHandler(this.EventHandlers)

	fact.Start(wait.NeverStop)
	return fact
}
