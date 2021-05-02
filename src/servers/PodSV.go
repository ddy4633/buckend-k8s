package servers

import (
	"fmt"
	"k8s-web/src/models"
	coreV1 "k8s.io/api/core/v1"
	"log"
)

type PodService struct {
	PodMap  *PodMap        `inject:"-"`
	Conmmon *CommonService `inject:"-"`
	Events  *EventMap      `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (p *PodService) ListPod(ns string) (res []*models.Pod) {
	pods, err := p.PodMap.List(ns)
	if err != nil {
		log.Println(err)
		//goft.Error(err)
		return make([]*models.Pod, 0)
	}
	for _, item := range pods {
		pod := &models.Pod{
			Name:       item.Name,
			NameSpace:  item.Namespace,
			Images:     p.Conmmon.GetContainerImages(item.Spec.Containers),
			NodeName:   item.Spec.NodeName,
			IP:         []string{p.Conmmon.TransfromPodTOString(item.Status.PodIP)},
			Phase:      fmt.Sprintf("%s", item.Status.Phase),
			IsRead:     p.getPodComplete(item),
			Message:    p.Events.GetMessages(item.Namespace, "Pod", item.Name),
			CreateTime: item.CreationTimestamp.String(),
			// 当第一次创建的时候返回的对象的len为零？
			RestartCount: p.getRestartCount(item),
			Age:          p.Conmmon.GetAge(item.CreationTimestamp.Time),
		}
		//fmt.Println(pod.Name, pod.IsRead)
		res = append(res, pod)
	}
	return res
}

// 判断pod是否完成
func (*PodService) getPodComplete(pod *coreV1.Pod) bool {
	for i, ava := range pod.Status.Conditions {
		//fmt.Println(pod.Name, i, ava)
		if string(ava.Status) != "True" {
			break
		} else if i == 3 && string(ava.Status) == "True" {
			return true
		}
	}
	return false
}

// 判断pod的可用状态
func (*PodService) getPodCondition(pod *coreV1.Pod) string {
	fmt.Println(pod.Status.Conditions, pod.Name)
	for _, ava := range pod.Status.Conditions {
		if string(ava.Status) != "True" {
			break
		}
	}
	return "not Available"
}

// 获取对象的重启次数
func (*PodService) getRestartCount(pod *coreV1.Pod) int32 {
	if len(pod.Status.ContainerStatuses) == 0 {
		return 0
	}
	return pod.Status.ContainerStatuses[0].RestartCount
}
