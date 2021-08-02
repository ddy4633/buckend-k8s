package config

import (
	"k8s-web/src/servers"

	"k8s.io/client-go/kubernetes"
)

type Maps struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewMaps() *Maps {
	return &Maps{}
}

// 初始化depolyment的
func (this *Maps) InitDeployMap() *servers.DeploymentMap {
	return &servers.DeploymentMap{}
}

// 初始化pod的
func (this *Maps) InitPodMap() *servers.PodMap {
	return &servers.PodMap{}
}

// 初始化namespace的
func (this *Maps) InitNameSpaceMap() *servers.NamespaceMap {
	return &servers.NamespaceMap{}
}

// 初始化Events
func (this *Maps) InitEventsMap() *servers.EventMap {
	return &servers.EventMap{}
}

// 初始化Service
func (this *Maps) InitServiceMap() *servers.ServiceMap {
	return &servers.ServiceMap{}
}

// 初始化Ingress
func (this *Maps) InitIngressMap() *servers.IngressMap {
	return &servers.IngressMap{}
}

// 初始化Secret
func (this *Maps) InitSecretMap() *servers.SecretMap {
	return &servers.SecretMap{}
}

//初始化Endpoints
func (*Maps) InitEndPointsMap() *servers.EndPointsMap {
	return &servers.EndPointsMap{}
}