package servers

import (
	"k8s-web/src/models"

	"github.com/shenyisyn/goft-gin/goft"
)

type SVservice struct {
	ServiceMaps *ServiceMap    `inject:"-"`
	Conmmon     *CommonService `inject:"-"`
}

func NewSVserver() *SVservice {
	return &SVservice{}
}

func (s *SVservice) ListAll(ns string) (res []*models.Services) {
	obj, err := s.ServiceMaps.List(ns)
	goft.Error(err)
	for _, item := range obj {
		list := &models.Services{
			Name:          item.Name,
			NameSpace:     item.Namespace,
			NetworkType:   item.Spec.Type,
			Address:       item.Spec.ClusterIP,
			ExportPort:    item.Spec.Ports[0].Port,
			ContainerPort: item.Spec.Ports[0].TargetPort,
			Protocol:      item.Spec.Ports[0].Protocol,
			Age:           s.Conmmon.GetAge(item.ObjectMeta.CreationTimestamp.Time),
			CreateTime:    item.ObjectMeta.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
		res = append(res, list)
	}
	return res
}
