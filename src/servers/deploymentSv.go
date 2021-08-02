package servers

import (
	"k8s-web/src/models"
	"sort"

	"github.com/shenyisyn/goft-gin/goft"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/api/networking/v1beta1"
)

type DeploymentService struct {
	DepMap  *DeploymentMap `inject:"-"`
	Conmmon *CommonService `inject:"-"`
}

func NewDeploymentService() *DeploymentService {
	return &DeploymentService{}
}

// 获取所有的deply信息
func (this *DeploymentService) ListAll(ns string) (res []*models.Deployment) {
	deplist, err := this.DepMap.List(ns)
	goft.Error(err)
	// 遍历deployment集合
	for _, item := range deplist {
		rep := &models.Deployment{
			Name:       item.Name,
			NameSpace:  item.Namespace,
			Replicas:   [3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas},
			IsComplete: this.getDeploymentsComplete(item),
			Message:    this.getDeploymentCondition(item),
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Images:     this.Conmmon.GetDeployMentImage(item),
			Age:        this.Conmmon.GetAge(item.CreationTimestamp.Time),
			Lables: 	item.Labels,
		}
		// 加入到副本集中
		res = append(res, rep)
	}
	return res
}

// 获取所有ingress信息
func (ing *IngressService) GetALLIngresss(ns string) []*models.Ingresses {
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

//// 获取指定的NS下的Deployment信息
//func (this *DeploymentService) ListDeployNs(ns string) *models.Deployment {
//	deploy, err := this.DepMap.
//}

// 判断deploy是否完成
func (*DeploymentService) getDeploymentsComplete(deploy *v1.Deployment) bool {
	return deploy.Status.Replicas == deploy.Status.AvailableReplicas
}

// 判断deploy可用状态信息
func (*DeploymentService) getDeploymentCondition(deploy *v1.Deployment) string {
	for _, item := range deploy.Status.Conditions {
		if string(item.Type) == "Available" && string(item.Status) != "True" {
			return item.Message
		}
	}
	return "not Available"
}
