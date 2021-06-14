package servers

import (
	"context"
	"k8s-web/src/models"
	"k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"sort"
	"strconv"
)

type IngressService struct {
	IngressMaps *IngressMap `inject:"-"`
	K8sClient *kubernetes.Clientset `inject:"-"`
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
		return make([]*models.Ingresses, 0)
	}
}

// 创建对应的ingress
func (ing *IngressService) CreateIngress(ingress *models.IngressCreate) error {
	// 组装Ingress的annotation
	annotations := make(map[string]string, 3)
	for _, va := range ingress.AnnotationsData[0].Rules {
		annotations[va.Key] = va.Value
	}
	className := annotations["nginx.ingress.kubernetes.io/ingress.class"]
	// 组装rule规则
	ingressRule := []v1beta1.IngressRule{}
	for _, ob := range ingress.Rules {
		httpRuleValue := &v1beta1.HTTPIngressRuleValue{}
		rulePaths := make([]v1beta1.HTTPIngressPath, 0)
		for _, Rulecfg := range ob.Paths {
			prot, err := strconv.Atoi(Rulecfg.Port)
			if err != nil {
				return err
			}
			rulePaths = append(rulePaths, v1beta1.HTTPIngressPath{
				Path: Rulecfg.Path,
				Backend: v1beta1.IngressBackend{
					ServiceName: Rulecfg.SvcName,
					ServicePort: intstr.FromInt(prot),
				},
			})
		}
		httpRuleValue.Paths = rulePaths
		rules := v1beta1.IngressRule{
			Host: ob.Host,
			IngressRuleValue: v1beta1.IngressRuleValue{
				HTTP: httpRuleValue,
			},
		}
		ingressRule = append(ingressRule, rules)
	}
	// 新建一个ingress的对象
	Newingress := &v1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "networking.k8s.io/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        ingress.Name,
			Namespace:   ingress.NameSpace,
			Annotations: annotations,
		},
		Spec: v1beta1.IngressSpec{
			IngressClassName: &className,
			Rules:            ingressRule,
		},
	}
	// 调用K8s-api创建Ingress对象
	_,err := ing.K8sClient.NetworkingV1beta1().Ingresses(ingress.NameSpace).
		Create(context.TODO(),Newingress,metav1.CreateOptions{})
	return err
}
