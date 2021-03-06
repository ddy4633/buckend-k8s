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

func (*ServiceConfig) Helper() *servers.Helper {
	return servers.NewHelper()
}

func (*ServiceConfig) IngressService() *servers.IngressService {
	return servers.NewIngressService()
}

func (*ServiceConfig) SecretService() *servers.SecretService {
	return servers.NewSecretService()
}

func (*ServiceConfig) EndPointsService() *servers.EndPointSV {
	return servers.NewEndPointSV()
}