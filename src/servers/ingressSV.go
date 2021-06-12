package servers

import (
	"k8s-web/src/models"
	"k8s.io/api/networking/v1beta1"
	"sort"
)

type IngressService struct {
	IngressMaps *IngressMap `inject:"-"`
}

func NewIngressService() *IngressService {
	return &IngressService{}
}

// 获取所有ingress信息
func (ing *IngressService) GetALLIngress(ns string) []*models.Ingresses {
	if va, ok := ing.IngressMaps.data.Load(ns); ok {
		obj := va.([]*v1beta1.Ingress)
		sort.Sort(v1beta1Ingress(obj))
		result := make([]*models.Ingresses, len(obj))
		for i, item := range obj {
			result[i] = &models.Ingresses{
				Name:       item.Name,
				NameSpace:  item.Namespace,
				CreateTime: item.CreationTimestamp.String(),
				Labels:     item.Labels,
				Status:     item.Status.String(),
				Rules:      item.Spec.Rules,
				Address:    item.Status.LoadBalancer.Ingress,
			}
		}
		return result
	} else {
		return make([]*models.Ingresses ,0)
	}
}

