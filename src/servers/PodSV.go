package servers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s-web/src/models"
	coreV1 "k8s.io/api/core/v1"
	"log"
)

type PodService struct {
	PodMap  *PodMap        `inject:"-"`
	Conmmon *CommonService `inject:"-"`
	Events  *EventMap      `inject:"-"`
	Helper  *Helper        `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

// 分页输出POD列表
func (p *PodService) PagePods(ns string, page, size int) *Paging {
	pods := p.ListPod(ns)
	readycount, allcount, ipods := p.GetPodtotle(pods)
	return p.Helper.PageResource(page, size, ipods).
		SetEXT(gin.H{"ReadyNum": readycount, "AllNum": allcount})
}

// 获取所有的Pod信息
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
			IP:         []string{p.Conmmon.TransfromPodTOString(item.Status.PodIP), item.Status.HostIP},
			Phase:      fmt.Sprintf("%s", item.Status.Phase),
			IsRead:     p.Conmmon.PodIsReady(item),
			Message:    p.Events.GetMessages(item.Namespace, "Pod", item.Name),
			CreateTime: item.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
			// 当第一次创建的时候返回的对象的len为零？
			Labels:         item.Labels,
			RestartCount:   p.getRestartCount(item),
			Age:            p.Conmmon.GetAge(item.CreationTimestamp.Time),
			ContainersName: p.GetContainers(item.Namespace, item.Name),
			Annotation:     item.Annotations,
			Tolerations:    item.Spec.Tolerations,
			Secret:         item.Spec.ServiceAccountName,
		}
		//fmt.Println(pod.Name, pod.IsRead)
		res = append(res, pod)
	}
	return res
}

// 获取pod的就绪总数和总数量
func (p *PodService) GetPodtotle(res []*models.Pod) (readyCount, allCount int, iPods []interface{}) {
	iPods = make([]interface{}, len(res))
	for i, pod := range res {
		allCount++
		iPods[i] = pod
		if pod.IsRead {
			readyCount++
		}
	}
	return
}

// 获取Container容器的列表
func (p *PodService) GetContainers(ns, podname string) []*models.Containers {
	result := make([]*models.Containers, 0)
	pods, err := p.PodMap.GetPod(ns, podname)
	if err == nil {
		for _, c := range pods.Spec.Containers {
			result = append(result, &models.Containers{
				Name: c.Name,
			})
		}
	}
	return result
}

// 判断pod是否完成
func (*PodService) getPodComplete(pod *coreV1.Pod) bool {
	return pod.Status.Phase == "Running"
}

// 判断pod的可用状态
func (*PodService) getPodCondition(pod *coreV1.Pod) string {
	for _, ava := range pod.Status.Conditions {
		if string(ava.Type) == "Available" && string(ava.Status) != "True" {
			return ava.Message
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
