package core

import (
	"fmt"
	"reflect"
	"sync"

	corev1 "k8s.io/api/core/v1"

	v1 "k8s.io/api/apps/v1"
)

// 对deployment的集合进行定义
type DeploymentMap struct {
	data sync.Map
}

// 新增对象
func (this *DeploymentMap) Add(deploy *v1.Deployment) {
	// 查询安全map下是否有名称为NS的deply没有则直接加入，有则组装后添加
	if list, ok := this.data.Load(deploy.Namespace); ok {
		list = append(list.([]*v1.Deployment), deploy)
		this.data.Store(deploy.Namespace, list)
	} else {
		this.data.Store(deploy.Namespace, []*v1.Deployment{deploy})
	}
}

// 删除对象
func (this *DeploymentMap) Delete(deploy *v1.Deployment) {
	// 若存在则循环遍历,判断要删除的是否存在列表如果存在则新建列表追加匹配到的0到前一项+匹配到的+1后续项
	if list, ok := this.data.Load(deploy.Namespace); ok {
		for i, item_dep := range list.([]v1.Deployment) {
			if item_dep.Name == deploy.Name {
				newlist := append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				this.data.Store(deploy.Namespace, newlist)
				break
			}
		}
	}
}

// 更新对象
func (this *DeploymentMap) Update(deploy *v1.Deployment) error {
	// 若存在则循环，匹配名称相等取出对应位置的deploy赋值对象
	if list, ok := this.data.Load(deploy.Namespace); ok {
		for i, item_deloy := range list.([]*v1.Deployment) {
			if item_deloy.Name == deploy.Name {
				list.([]*v1.Deployment)[i] = deploy
				break
			}
		}
	} else {
		return fmt.Errorf("not found the Deplyment ", deploy.Name)
	}
	return nil
}

// 查询全部的对象
func (this *DeploymentMap) List(ns string) ([]*v1.Deployment, error) {
	if list, ok := this.data.Load(ns); ok {
		return list.([]*v1.Deployment), nil
	}
	return nil, fmt.Errorf("record not foound deployment Object")
}

// 查询单个对象
func (this *DeploymentMap) GetDeploy(ns, name string) (*v1.Deployment, error) {
	if list, ok := this.data.Load(ns); ok {
		for _, item := range list.([]*v1.Deployment) {
			if item.Name == name {
				return item, nil
			}
		}
	}
	return nil, fmt.Errorf("Not found this Deployment", ns, "/", name)
}

type PodMap struct {
	data sync.Map
}

// 新增
func (p *PodMap) Add(pod *corev1.Pod) {
	if list, ok := p.data.Load(pod.Namespace); ok {
		list = append(list.([]*corev1.Pod), pod)
		p.data.Store(pod.Namespace, list)
	} else {
		p.data.Store(pod.Namespace, []*corev1.Pod{pod})
	}
}

// 更新
func (p *PodMap) Update(pod *corev1.Pod) error {
	if list, ok := p.data.Load(pod.Namespace); ok {
		for i, item := range list.([]*corev1.Pod) {
			if item.Name == pod.Name {
				list.([]*corev1.Pod)[i] = pod
				break
			}
		}
	} else {
		return fmt.Errorf("not found the pod ", pod.Name)
	}
	return nil
}

// 删除
func (p *PodMap) Delete(pod *corev1.Pod) {
	if list, ok := p.data.Load(pod.Name); ok {
		for i, item := range list.([]*corev1.Pod) {
			if item.Name == pod.Name {
				newlist := append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[i+1:]...)
				p.data.Store(pod.Namespace, newlist)
			}
		}
	}
}

// 查询所有
func (p *PodMap) List(ns string) ([]*corev1.Pod, error) {
	res := make([]*corev1.Pod, 0)
	if list, ok := p.data.Load(ns); ok {
		for _, pod := range list.([]*corev1.Pod) {
			res = append(res, pod)
		}
		return res, nil
	}
	return res, fmt.Errorf("record not foound Pod Object")
}

// 查询单个
func (p *PodMap) GetPod(ns, name string) (*corev1.Pod, error) {
	if list, ok := p.data.Load(ns); ok {
		for _, item := range list.([]*corev1.Pod) {
			if item.Name == name {
				return item, nil
			}
		}
	}
	return nil, fmt.Errorf("Not found this Pod", ns, "/", name)
}

// 标签获取列表
func (p *PodMap) LabelsList(ns string, labels []map[string]string) ([]*corev1.Pod, error) {
	res := make([]*corev1.Pod, 0)
	if list, ok := p.data.Load(ns); ok {
		for _, item := range list.([]*corev1.Pod) {
			for _, label := range labels {
				// 对切片，map，结构体进行深度的对比
				if reflect.DeepEqual(item.Labels, label) {
					res = append(res, item)
				}
			}
		}
		return res, nil
	}
	return nil, fmt.Errorf("pods not found ")
}

// debug Get pod
func (p *PodMap) DEBUG_ListByNS(ns string) []*corev1.Pod {
	ret := make([]*corev1.Pod, 0)
	if list, ok := p.data.Load(ns); ok {
		for _, pod := range list.([]*corev1.Pod) {
			ret = append(ret, pod)
		}

	}
	return ret
}
