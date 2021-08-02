package servers

import (
	"fmt"
	"k8s-web/src/models"
	"k8s.io/api/networking/v1beta1"
	"reflect"
	"sort"
	"sync"

	corev1 "k8s.io/api/core/v1"

	v1 "k8s.io/api/apps/v1"
)

type MapItems []*MapItem

// 公共模型
type MapItem struct {
	key   string
	value interface{}
}

// 把sync.map的数据转化为切片提供给后续的处理
func convertItems(m sync.Map) MapItems {
	// 初始化
	items := make(MapItems, 0)
	m.Range(func(key interface{}, value interface{}) bool {
		items = append(items, &MapItem{key: key.(string), value: value})
		return true
	})
	return items
}

// 求长度
func (m MapItems) Len() int {
	return len(m)
}

// 正排序
func (m MapItems) Less(i, j int) bool {
	return m[i].key < m[j].key
}

// 交换位置
func (m MapItems) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 对deployment的集合进行定义
type DeploymentMap struct {
	data sync.Map
}

// Deploymen新增对象
func (this *DeploymentMap) Add(deploy *v1.Deployment) {
	// 查询安全map下是否有名称为NS的deply没有则直接加入，有则组装后添加
	if list, ok := this.data.Load(deploy.Namespace); ok {
		list = append(list.([]*v1.Deployment), deploy)
		this.data.Store(deploy.Namespace, list)
	} else {
		this.data.Store(deploy.Namespace, []*v1.Deployment{deploy})
	}
}

// Deploymen删除对象
func (this *DeploymentMap) Delete(deploy *v1.Deployment) {
	// 若存在则循环遍历,判断要删除的是否存在列表如果存在则新建列表追加匹配到的0到前一项+匹配到的+1后续项
	if list, ok := this.data.Load(deploy.Namespace); ok {
		for i, item_dep := range list.([]*v1.Deployment) {
			if item_dep.Name == deploy.Name {
				newlist := append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				this.data.Store(deploy.Namespace, newlist)
				break
			}
		}
	}
}

// Deploymen更新对象
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

// Deploymen查询全部的对象
func (this *DeploymentMap) List(ns string) ([]*v1.Deployment, error) {
	if list, ok := this.data.Load(ns); ok {
		return list.([]*v1.Deployment), nil
	}
	return nil, fmt.Errorf("record not foound deployment Object")
}

// Deploymen查询单个对象
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

// 对Service集合进行定义
type ServiceMap struct {
	data sync.Map
}

// Service新增
func (s *ServiceMap) Add(sv *corev1.Service) {
	if server, ok := s.data.Load(sv.Namespace); ok {
		server := append(server.([]*corev1.Service), sv)
		s.data.Store(sv.Namespace, server)
	} else {
		s.data.Store(sv.Namespace, []*corev1.Service{sv})
	}
}

// Service变更
func (s *ServiceMap) Update(sv *corev1.Service) {
	if server, ok := s.data.Load(sv.Namespace); ok {
		for i, va := range server.([]*corev1.Service) {
			if va.Name == sv.Name {
				server.([]*corev1.Service)[i] = sv
				break
			}
		}
	}
}

// Service删除
func (s *ServiceMap) Delete(sv *corev1.Service) {
	if server, ok := s.data.Load(sv.Namespace); ok {
		for i, va := range server.([]*corev1.Service) {
			if sv.Name == va.Name {
				list := append(server.([]*corev1.Service)[:i], server.([]*corev1.Service)[:i+1]...)
				s.data.Store(sv.Namespace, list)
			}
		}
	}
}

// Service查询所有
func (s *ServiceMap) List(ns string) ([]*corev1.Service, error) {
	list := make([]*corev1.Service, 0)
	if sv, ok := s.data.Load(ns); ok {
		for _, va := range sv.([]*corev1.Service) {
			list = append(list, va)
		}
		return list, nil
	}
	return list, fmt.Errorf("record not foound Services Object")
}

// 对Pod进行排序
type CorePods []*corev1.Pod

// 求本身的长度
func (c CorePods) Len() int {
	return len(c)
}

// 利用时间来做(正排序)
func (c CorePods) Less(i, j int) bool {
	return c[i].CreationTimestamp.Time.Before(c[j].CreationTimestamp.Time)
}

// 数据交换实现
func (c CorePods) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// 对Pod的集合进行定义
type PodMap struct {
	data sync.Map
}

// Pod新增
func (p *PodMap) Add(pod *corev1.Pod) {
	if list, ok := p.data.Load(pod.Namespace); ok {
		list = append(list.([]*corev1.Pod), pod)
		p.data.Store(pod.Namespace, list)
	} else {
		p.data.Store(pod.Namespace, []*corev1.Pod{pod})
	}
}

// Pod更新
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

// Pod删除
func (p *PodMap) Delete(pod *corev1.Pod) {
	if list, ok := p.data.Load(pod.Namespace); ok {
		for i, item := range list.([]*corev1.Pod) {
			if item.Name == pod.Name {
				newlist := append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[i+1:]...)
				p.data.Store(pod.Namespace, newlist)
			}
		}
	}
}

// Pod查询所有
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

// Pod查询单个
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

// Pod标签获取列表
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

// 返回排序后的pod列表
func (p *PodMap) ListByNS(ns string) []*corev1.Pod {
	if list, ok := p.data.Load(ns); ok {
		ret := list.([]*corev1.Pod)
		// 进行排序
		sort.Sort(CorePods(ret))
		return ret
	}
	return nil
}

// 对Namespace集合进行定义
type NamespaceMap struct {
	data sync.Map
}

// 新增namespace
func (ns *NamespaceMap) Add(namespace *corev1.Namespace) {
	ns.data.Store(namespace.Name, namespace)
}

// 更新namespace
func (ns *NamespaceMap) Update(namespace *corev1.Namespace) {
	ns.data.Store(namespace.Name, namespace)
}

// 删除namespace
func (ns *NamespaceMap) Delete(namespace *corev1.Namespace) {
	ns.data.Delete(namespace.Name)
}

// 查询单个ns
func (ns *NamespaceMap) Get(namespace *corev1.Namespace) *corev1.Namespace {
	if va, ok := ns.data.Load(namespace); ok {
		return va.(*corev1.Namespace)
	}
	return nil
}

// 返回所有的Namespace
func (ns *NamespaceMap) ListAll() []*models.NsModel {
	// 返回封装好的对象
	item := convertItems(ns.data)
	ret := make([]*models.NsModel, len(item))
	// 进行正向的排序
	sort.Sort(item)
	// 循环排序后的对象封装好对象返回
	for index, value := range item {
		ret[index] = &models.NsModel{Name: value.key}
	}
	/*			原始的方案进行循环拼接数据
	ns.data.Range(func(key, value interface{}) bool {
		ret = append(ret, &models.NsModel{Name: key.(string)})
		return true
	}) */
	return ret
}

// 对Event事件进行底层定义
type EventMap struct {
	data sync.Map
}

// 获取某个事件的信息
func (e *EventMap) GetMessages(ns, kind, name string) string {
	key := fmt.Sprintf("%s_%s_%s", ns, kind, name)
	if v, ok := e.data.Load(key); ok {
		return v.(*corev1.Event).Message
	}
	return ""
}

// 对ingress资源对象进行存储
type IngressMap struct {
	data sync.Map
}

// 存储数据
func (ing *IngressMap) Add(ingress *v1beta1.Ingress) {
	if va, ok := ing.data.Load(ingress.Namespace); ok {
		list := append(va.([]*v1beta1.Ingress), ingress)
		ing.data.Store(ingress.Namespace, list)
	} else {
		ing.data.Store(ingress.Namespace, []*v1beta1.Ingress{ingress})
	}
}

// 数据更新
func (ing *IngressMap) Update(ingress *v1beta1.Ingress) error {
	if va, ok := ing.data.Load(ingress.Namespace); ok {
		for i, ingress_obj := range va.([]*v1beta1.Ingress) {
			if ingress_obj.Name == ingress.Name {
				va.([]*v1beta1.Ingress)[i] = ingress
			}
		}
		return nil
	}
	return fmt.Errorf("%s - ingress-%s is not found", ingress.Namespace, ingress.Name)
}

// 删除数据
func (ing *IngressMap) Delete(ingress *v1beta1.Ingress) {
	if va, ok := ing.data.Load(ingress.Namespace); ok {
		for i, ingress_obj := range va.([]*v1beta1.Ingress) {
			if ingress_obj.Name == ingress.Name {
				list := append(va.([]*v1beta1.Ingress)[:i], va.([]*v1beta1.Ingress)[i+1:]...)
				ing.data.Store(ingress_obj.Namespace, list)
				break
			}
		}
	}
}

// 获取单个ingress信息
func (ing *IngressMap) GetIngress(ns, name string) *v1beta1.Ingress {
	if va, ok := ing.data.Load(ns); ok {
		for _, ingress := range va.([]*v1beta1.Ingress) {
			if ingress.Name == name {
				return ingress
			}
		}
	}
	return nil
}


// 获取全部的ingress信息
func (ing *IngressMap) ListIngress(ns string) ([]*v1beta1.Ingress, error) {
	if va, ok := ing.data.Load(ns); ok {
		return va.([]*v1beta1.Ingress), nil
	}
	return nil, fmt.Errorf("%s - ingress is not found", ns)
}

// ingress对象排序的实现
type v1beta1Ingress []*v1beta1.Ingress

func (ing v1beta1Ingress) Len() int {
	return len(ing)
}

func (ing v1beta1Ingress) Less(i, j int) bool {
	// 根据时间间隔（倒排序）
	return ing[i].CreationTimestamp.Time.After(ing[j].CreationTimestamp.Time)
}

func (ing v1beta1Ingress) Swap(i, j int) {
	ing[i], ing[j] = ing[j], ing[i]
}

// secret底层存储
type SecretMap struct {
	data sync.Map
}

func (se *SecretMap) Add(secret *corev1.Secret) {
	if va,ok := se.data.Load(secret.Namespace);ok{
		list := append(va.([]*corev1.Secret),secret)
		se.data.Store(secret.Namespace,list)
	}else {
		se.data.Store(secret.Namespace,[]*corev1.Secret{secret})
	}
}

func (se *SecretMap) Delete(secret *corev1.Secret) {
	if va,ok := se.data.Load(secret.Namespace);ok{
		for k,value := range va.([]*corev1.Secret) {
			if value.Name == secret.Name {
				list := append(va.([]*corev1.Secret)[:k],va.([]*corev1.Secret)[:k+1]...)
				se.data.Store(secret.Namespace,list)
				break
			}
		}
	}
}

func (se *SecretMap) Update(secret *corev1.Secret) error {
	if va,ok := se.data.Load(secret.Namespace);ok{
		for k,value := range va.([]*corev1.Secret) {
			if value.Name == secret.Name {
				va.([]*corev1.Secret)[k]=secret
			}
		}
		return nil
	}else {
		return fmt.Errorf("not found %s/%s",secret.Namespace,secret.Name)
	}
}

func (se *SecretMap) Get(ns string,name string) *corev1.Secret{
	if va,ok := se.data.Load(ns);!ok {
		for _,v := range va.([]*corev1.Secret){
			if v.Name == name {
				return v
			}
		}
	}
	return nil
}

func (se *SecretMap) GetALL(ns string) ([]*corev1.Secret,error) {
	if va,ok := se.data.Load(ns);ok {
		return va.([]*corev1.Secret),nil
	}else {
		return nil,fmt.Errorf("not found secret in %s",ns)
	}
}

// 嗯，这次是endpoints信息
type EndPointsMap struct {
	data sync.Map
}

func (ep *EndPointsMap) Add(eps *corev1.Endpoints) {
	if va,ok := ep.data.Load(eps.Namespace);ok {
		list := append(va.([]*corev1.Endpoints),eps)
		ep.data.Store(eps.Namespace,list)
	}else{
		ep.data.Store(eps.Namespace,eps)
	}
}
func (ep *EndPointsMap) Delete(eps *corev1.Endpoints) {
	if va,ok:= ep.data.Load(eps.Namespace);ok {
		for i,item := range va.([]*corev1.Endpoints) {
			if item.Name == eps.Name {
				newep := append(va.([]*corev1.Endpoints)[:i],va.([]*corev1.Endpoints)[i-1:]...)
				ep.data.Store(eps.Namespace,newep)
				break
			}
		}
	}
}
func (ep *EndPointsMap) update(eps *corev1.Endpoints) {
	if va,ok := ep.data.Load(eps.Namespace);ok {
		for i,item := range va.([]*corev1.Endpoints) {
			if item.Name == eps.Name {
				va.([]*corev1.Endpoints)[i]=eps
			}
		}
	}
}

func (ep *EndPointsMap) Getns(ns string) ([]*corev1.Endpoints,error) {
	if va,ok := ep.data.Load(ns);ok{
		fmt.Printf("%v",va.([]*corev1.Endpoints))
		return va.([]*corev1.Endpoints),nil
	}else {
		return nil,fmt.Errorf("not found EndPonits in %s",ns)
	}
}

func (ep *EndPointsMap) Get(ns,name string) (*corev1.Endpoints,error) {
	va,_ := ep.data.Load(ns)
	for i,item := range va.([]*corev1.Endpoints) {
		if item.Name == name {
			return va.([]*corev1.Endpoints)[i],nil
		}
	}
	return nil,fmt.Errorf("not found EndPonits/%s in %s",name,ns)
}

// 存储Node节点信息
type NodeMap struct {
	data sync.Map
}

func (nd *NodeMap) Add(node *corev1.Node) {
	nd.data.Store(node.Name,node)
}

func (nd *NodeMap) Delete(node *corev1.Node) {
	nd.data.Delete(node.Name)
}

func (nd *NodeMap) Update(node *corev1.Node) {
	nd.data.Store(node.Name,node)
}

func (nd *NodeMap) Get(name string) *corev1.Node {
	va,_ := nd.data.Load(name)
	return va.(*corev1.Node)
}

func (nd *NodeMap) GetAll() []*corev1.Node {
	result := []*corev1.Node{}
	nd.data.Range(func(key, value interface{}) bool {
		result = append(result,value.(*corev1.Node))
		return true
	})
	return result
}