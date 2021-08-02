package servers

import (
	"k8s-web/src/models"
	corev1 "k8s.io/api/core/v1"
	"log"
)

type EndPointSV struct {
	endPointsMap *EndPointsMap `inject:"-"`
}

func NewEndPointSV() *EndPointSV {
	return &EndPointSV{}
}

func (ep *EndPointSV) Getall(ns string) (res []*models.EndPoints) {
	eplist, err := ep.endPointsMap.Getns(ns)
	log.Println(err)
	for _, endpoint := range eplist {
		status,port := ep.getAdress(endpoint)
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
func (ep *EndPointSV) getAdress(endpoints *corev1.Endpoints) (models.EndpointsStatus,models.EndPonitsPort) {
	var adress, noreadyADR,node,hostname []string
	for _, sub := range endpoints.Subsets {
		for _, oneadress := range sub.Addresses {
			adress = append(adress, oneadress.IP)
			node = append(node,*oneadress.NodeName)
			hostname = append(hostname,oneadress.Hostname)
		}
		for _, notIP := range sub.NotReadyAddresses {
			noreadyADR = append(noreadyADR, notIP.IP)
		}
	}
	return models.EndpointsStatus{
		Addresss:          adress,
		NotReadyAddresses: noreadyADR,
		TargetRefName: endpoints.Subsets[0].Addresses[0].TargetRef.Name,
		NodeName: node,
		HostName: hostname,
	},models.EndPonitsPort{
			Name: endpoints.Subsets[0].Ports[0].Name,
			Port: string(endpoints.Subsets[0].Ports[0].Port),
			Protocol: endpoints.Subsets[0].Ports[0].Protocol,
		}
}

