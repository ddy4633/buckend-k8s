package servers

import (
	"fmt"
	"time"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

// 获取deployment的镜像名称
func (this *CommonService) GetDeployMentImage(deploy *v1.Deployment) string {
	return this.GetContainerImages(deploy.Spec.Template.Spec.Containers)
}

// 获取Pod中某个Container镜像的名称
func (*CommonService) GetContainerImages(containers []corev1.Container) string {
	images := containers[0].Image
	if len(containers) > 1 {
		images = fmt.Sprintf("%s+其他%d个镜像", images, len(containers)-1)
	}
	return images
}

// 类型转换为string
func (*CommonService) TransfromPodTOString(obj interface{}) string {
	return obj.(string)
}

/* 			新版获取POD是否正常的判断方式
PodScheduled 	-> Pod已经被调度到指定的节点
ContainersReady -> Pod中所有的容器已经处于就绪状态
Initialized		-> 所有的Init容器都已经成功的启动
Ready			-> Pod可以为请求提供服务
*/
func (*CommonService) PodIsReady(pod *corev1.Pod) bool {
	// 如果不为run则返回false
	if pod.Status.Phase != "Running" {
		return false
	}
	// 循环判断一个Pod中的所有container是否都是True，和当前pod的服务状态是否active
	for _, Rcon := range pod.Spec.ReadinessGates {
		for _, condition := range pod.Status.Conditions {
			if condition.Type == Rcon.ConditionType && condition.Status != "True" {
				return false
			}
		}
	}
	for _, condition := range pod.Status.Conditions {
		if condition.Status != "True" {
			return false
		}
	}
	return true
}

// 获取工作了多少时间时间
func (*CommonService) GetAge(t time.Time) string {
	t1 := time.Now()
	sub := t1.Sub(t)
	if hours := sub.Hours(); hours > 0 {
		return fmt.Sprintf("%.0f", hours/24)
	} else {
		return "0"
	}
}

