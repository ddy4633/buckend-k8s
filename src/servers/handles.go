package servers

import (
	"fmt"
	"k8s-web/src/models"
	"k8s-web/src/wscore"
	"k8s.io/api/networking/v1beta1"
	"log"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// DeployMent回调的方法
type DeployHandler struct {
	DepMap     *DeploymentMap     `inject:"-"`
	DeployMent *DeploymentService `inject:"-"`
}

func (this *DeployHandler) OnAdd(obj interface{}) {
	this.DepMap.Add(obj.(*v1.Deployment))
	returnMsg("deployment",
		obj.(*v1.Deployment).Namespace,
		this.DeployMent.ListAll(obj.(*v1.Deployment).Namespace))
}

func (this *DeployHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	if err := this.DepMap.Update(newObj.(*v1.Deployment)); err != nil {
		log.Println(err)
	} else {
		returnMsg("deployment",
			newObj.(*v1.Deployment).Namespace,
			this.DeployMent.ListAll(newObj.(*v1.Deployment).Namespace))
	}
}

func (this *DeployHandler) OnDelete(obj interface{}) {
	if objs, ok := obj.(*v1.Deployment); ok {
		this.DepMap.Delete(objs)
		returnMsg("deployment",
			obj.(*v1.Deployment).Namespace,
			this.DeployMent.ListAll(obj.(*v1.Deployment).Namespace))
	}
}

// Pod的对象回调方法
type PodHandler struct {
	PodMaps     *PodMap     `inject:"-"`
	PodServices *PodService `inject:"-"`
}

func (p *PodHandler) OnAdd(obj interface{}) {
	p.PodMaps.Add(obj.(*corev1.Pod))
	// 多层的封装
	returnMsg("pod",
		obj.(*corev1.Pod).Namespace,
		p.PodServices.PagePods(obj.(*corev1.Pod).Namespace, 1, 8))
}

func (p *PodHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	if err := p.PodMaps.Update(newObj.(*corev1.Pod)); err != nil {
		log.Println(err)
	} else {
		returnMsg("pod",
			newObj.(*corev1.Pod).Namespace,
			p.PodServices.PagePods(newObj.(*corev1.Pod).Namespace, 1, 8))
	}
}

func (p *PodHandler) OnDelete(obj interface{}) {
	if objs, ok := obj.(*corev1.Pod); ok {
		p.PodMaps.Delete(objs)
		returnMsg("pod",
			obj.(*corev1.Pod).Namespace,
			p.PodServices.PagePods(obj.(*corev1.Pod).Namespace, 1, 8))
	}
}

// NameSpace的对象回调方法
type NameSpaceHandler struct {
	NsMap *NamespaceMap `inject:"-"`
}

func (ns *NameSpaceHandler) OnAdd(obj interface{}) {
	ns.NsMap.Add(obj.(*corev1.Namespace))
}

func (ns *NameSpaceHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	ns.NsMap.Update(newObj.(*corev1.Namespace))
}

func (ns *NameSpaceHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*corev1.Namespace); ok {
		ns.NsMap.Delete(d)
	}
}

// Event相关的Handler
type EventHandler struct {
	EventMap *EventMap `inject:"-"`
}

// 处理存储事件
func (e *EventHandler) StoreData(obj interface{}, isdelete bool) {
	if event, ok := obj.(*corev1.Event); ok {
		key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)
		if !isdelete {
			e.EventMap.data.Store(key, event)
		} else {
			e.EventMap.data.Delete(key)
		}
	}
}

func (e *EventHandler) OnAdd(obj interface{}) {
	e.StoreData(obj, false)
}

func (e *EventHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	e.StoreData(newObj, false)
}

func (e *EventHandler) OnDelete(obj interface{}) {
	e.StoreData(obj, true)
}

// 公共返回的消息格式方法
func returnMsg(kind, ns string, data interface{}) {
	msg := &models.ReturnMsg{
		Type: kind,
		Ns:   ns,
		Data: data,
	}
	wscore.ClientMap.Sendall(msg)
}

// Service相关的Handler
type ServiceHandler struct {
	ServiceMap *ServiceMap `inject:"-"`
}

func (s *ServiceHandler) OnAdd(obj interface{}) {
	s.ServiceMap.Add(obj.(*corev1.Service))
}

func (s *ServiceHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	s.ServiceMap.Update(newObj.(*corev1.Service))
}

func (s *ServiceHandler) OnDelete(obj interface{}) {
	s.ServiceMap.Delete(obj.(*corev1.Service))
}

// Ingress回调的对象
type IngressHandle struct {
	IngressMap     *IngressMap     `inject:"-"`
	IngressService *IngressService `inject:"-"`
}

func (ing *IngressHandle) OnAdd(obj interface{}) {
	ing.IngressMap.Add(obj.(*v1beta1.Ingress))
	returnMsg("ingress",
		obj.(*v1beta1.Ingress).Namespace,
		ing.IngressService.GetALLIngress(obj.(*v1beta1.Ingress).Namespace))
}

func (ing *IngressHandle) OnUpdate(lodObj, newObj interface{}) {
	ing.IngressMap.Update(newObj.(*v1beta1.Ingress))
	returnMsg("ingress",
		newObj.(*v1beta1.Ingress).Namespace,
		ing.IngressService.GetALLIngress(newObj.(*v1beta1.Ingress).Namespace))
}

func (ing *IngressHandle) OnDelete(obj interface{}) {
	ing.IngressMap.Delete(obj.(*v1beta1.Ingress))
	returnMsg("ingress",
		obj.(*v1beta1.Ingress).Namespace,
		ing.IngressService.GetALLIngress(obj.(*v1beta1.Ingress).Namespace))
}

// Secret资源对象回调
type SecretHandle struct {
	SecretMaps    *SecretMap     `inject:"-"`
	SecretService *SecretService `inject:"-"`
}

func (se *SecretHandle) OnAdd(obj interface{}) {
	se.SecretMaps.Add(obj.(*corev1.Secret))
	returnMsg("secret",
		obj.(*corev1.Secret).Namespace,
		se.SecretService.Getall(obj.(*corev1.Secret).Namespace))
}

func (se *SecretHandle) OnUpdate(lodObj, newObj interface{}) {
	se.SecretMaps.Update(newObj.(*corev1.Secret))
	returnMsg("secret",
		newObj.(*corev1.Secret).Namespace,
		se.SecretService.Getall(newObj.(*corev1.Secret).Namespace))
}

func (se *SecretHandle) OnDelete(obj interface{}) {
	se.SecretMaps.Delete(obj.(*corev1.Secret))
	returnMsg("secret",
		obj.(*corev1.Secret).Namespace,
		se.SecretService.Getall(obj.(*corev1.Secret).Namespace))
}
