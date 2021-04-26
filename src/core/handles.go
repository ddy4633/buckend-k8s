package core

import (
	"log"

	corev1 "k8s.io/api/core/v1"

	v1 "k8s.io/api/apps/v1"
)

// DeployMent回调的方法
type DeployHandler struct {
	DepMap *DeploymentMap `inject:"-"`
}

func (this *DeployHandler) OnAdd(obj interface{}) {
	this.DepMap.Add(obj.(*v1.Deployment))
}

func (this *DeployHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	if err := this.DepMap.Update(newObj.(*v1.Deployment)); err != nil {
		log.Println(err)
	}
}

func (this *DeployHandler) OnDelete(obj interface{}) {
	if obj, ok := obj.(*v1.Deployment); ok {
		this.DepMap.Delete(obj)
	}
}

// Pod的对象回调方法
type PodHandler struct {
	PodMap *PodMap `inject:"-"`
}

func (p *PodHandler) OnAdd(obj interface{}) {
	p.PodMap.Add(obj.(*corev1.Pod))
}

func (p *PodHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	if err := p.PodMap.Update(newObj.(*corev1.Pod)); err != nil {
		log.Println(err)
	}
}

func (p *PodHandler) OnDelete(obj interface{}) {
	if obj, ok := obj.(*corev1.Pod); ok {
		p.PodMap.Delete(obj)
	}
}
