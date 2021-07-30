package servers

import "k8s-web/src/models"

type NodeService struct {
	NodeMaps NodeMap `inject:"-"`
}

func (nd *NodeService) Getall() (renode []*models.Nodes) {
	nodes := nd.NodeMaps.GetAll()
	for _, node := range nodes {
		obj := &models.Nodes{
			Name:       node.Name,
			Labels:     node.Labels,
			Annotation: node.Annotations,
			Taints:     node.Spec.Taints,
			//Caps: node.Status.Allocatable.Cpu(),
		}
		renode = append(renode, obj)
	}
	return nil
}
