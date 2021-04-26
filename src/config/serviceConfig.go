package config

import "k8s-web/src/servers"

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (*ServiceConfig) CommonService() *servers.CommonService {
	return servers.NewCommonService()
}

func (*ServiceConfig) DeploymentService() *servers.DeploymentService {
	return servers.NewDeploymentService()
}

func (*ServiceConfig) PodService() *servers.PodService {
	return servers.NewPodService()
}

func (*ServiceConfig) NsService() *servers.NamespaceService {
	return servers.NewNamespaceService()
}

func (*ServiceConfig) SvService() *servers.SVservice {
	return servers.NewSVserver()
}
