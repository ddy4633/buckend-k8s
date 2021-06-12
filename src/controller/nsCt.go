package controller

import (
	"k8s-web/src/servers"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type NsCtl struct {
	NsService *servers.NamespaceMap `inject:"-"`
}

func NewNsCtl() *NsCtl {
	return &NsCtl{}
}

func (ns *NsCtl) ListAll(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": ns.NsService.ListAll(),
	}
}

func (ns *NsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/nslist", ns.ListAll)
}

func (*NsCtl) Name() string {
	return "NsCtl"
}
