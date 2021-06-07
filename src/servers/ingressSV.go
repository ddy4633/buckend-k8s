package servers

import (
	"fmt"
	"k8s-web/src/models"
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
	if va, err := ing.IngressMap.ListIngress(ns); err == nil {
		// 进行排序处理
		sort.Sort(v1beta1Ingress(va))
		// 初始化自定义模型
		result := make([]*models.Ingresses, len(va))
		for i, ingress := range va {
			result[i] = &models.Ingresses{
				Name:       ingress.Name,
				NameSpace:  ingress.Namespace,
				CreateTime: ingress.CreationTimestamp.String(),
				Labels:     ingress.Labels,
				Status:     ingress.Status.String(),
			}
		}
		return result
	} else {
		fmt.Println(err)
		return make([]*models.Ingresses, 0)
	}
}
