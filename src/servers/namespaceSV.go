package servers

import (
	"k8s-web/src/models"
)

type NamespaceService struct {
	NSMap *NamespaceMap `inject:"-"`
}

func NewNamespaceService() *NamespaceService {
	return &NamespaceService{}
}

func (ns *NamespaceService) ListAll() []*models.NsModel {
	return ns.NSMap.ListAll()
}
