package servers

import (
	"k8s-web/src/models"
	corev1 "k8s.io/api/core/v1"
	"log"
)

type EndPointSV struct {
	EndPointsMap *EndPointsMap `inject:"-"`
}

func NewEndPointSV() *EndPointSV {
	return &EndPointSV{}
}

func (ep *EndPointSV) Getall(ns string) (res []*models.EndPoints) {
	eplist, err := ep.EndPointsMap.Getns(ns)
	if err != nil {
		log.Println(err)
	}
	for _, endpoint := range eplist {
		status, port := ep.getAdress(endpoint)
		result := &models.EndPoints{
			Name:      endpoint.Name,
			NameSpace: endpoint.Namespace,
			Lables:    endpoint.Labels,
			CreatTime: endpoint.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Status:    status,
			Port:      port,
		}
		res = append(res, result)
	}
	return
}

// 获取单个endpoint下的adress信息
func (ep *EndPointSV) getAdress(endpoints *corev1.Endpoints) ( res []models.EndpointsStatus, resport []models.EndPonitsPort) {
	for _, sub := range endpoints.Subsets {
		status := models.EndpointsStatus{}
		epport := models.EndPonitsPort{}
		var address []string
		for _, oneadress := range sub.Addresses {
			address = append(address,oneadress.IP)
		}
		status.Addresss = address
		for _, notIP := range sub.NotReadyAddresses {
			status.NotReadyAddresses = notIP.IP
		}
		for _, p := range sub.Ports {
			epport.Name = p.Name
			epport.Port = p.Port
			epport.Protocol = p.Protocol
			resport = append(resport,epport)
		}
		res = append(res,status)
	}
	return res,resport
}
