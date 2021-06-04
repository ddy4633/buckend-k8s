package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-web/src/servers"
)

type IngressCtl struct {
	IngressService servers.IngressService `inject:"-"`
}

func NewIngressCtl() *IngressCtl {
	return &IngressCtl{}
}

func (*IngressCtl) Name() string {
	return "IngressCtl"
}

func (ing *IngressCtl) ListAll(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": ing.IngressService.GetALLIngress(c.DefaultQuery("ns","default")),
	}
}

func (ing *IngressCtl) Build(goft goft.Goft) {
	goft.Handle("GET","/v1/ingress",ing.ListAll)
}