package config

import "k8s-web/src/servers"

// 注入回调
type K8sHandle struct {
}

func NewK8sHandle() *K8sHandle {
	return &K8sHandle{}
}

func (this *K8sHandle) DeployHandlers() *servers.DeployHandler {
	return &servers.DeployHandler{}
}

func (this *K8sHandle) PodHandlers() *servers.PodHandler {
	return &servers.PodHandler{}
}

func (this *K8sHandle) NamespaceHandler() *servers.NameSpaceHandler {
	return &servers.NameSpaceHandler{}
}

func (this *K8sHandle) EventHandler() *servers.EventHandler {
	return &servers.EventHandler{}
}

func (this *K8sHandle) ServiceHandler() *servers.ServiceHandler {
	return &servers.ServiceHandler{}
}
