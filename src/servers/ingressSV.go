package servers

import (
	"k8s-web/src/models"
	"k8s.io/api/networking/v1beta1"
	"sort"
)

type IngressService struct {
	IngressMap *IngressMap
}

func NewIngressService() *IngressService {
	return &IngressService{}
}


// 获取所有ingress信息
func (ing *IngressService) GetALLIngress(ns string) []*models.Ingresses {
	if va,ok := ing.IngressMap.data.Load(ns);ok {
		newlist := va.([]*v1beta1.Ingress)
		// 进行排序处理
		sort.Sort(v1beta1Ingress(newlist))
		// 初始化自定义模型
		result := make([]*models.Ingresses,len(newlist))
		for i,ingress := range newlist {
			result[i]=&models.Ingresses{
				Name:       ingress.Name,
				NameSpace:  ingress.Namespace,
				CreateTime: ingress.CreationTimestamp.String(),
				Labels: ingress.Labels,
				Status: ingress.Status.String(),
			}
		}
		return result
	}
	return nil
}