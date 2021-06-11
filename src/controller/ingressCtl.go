package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-web/src/servers"
)

type IngressCtl struct {
	IngressService *servers.IngressService `inject:"-"`
	IngressMap servers.IngressMap `inject:"-"`
}

func NewIngressCtl() *IngressCtl {
	return &IngressCtl{}
}

func (*IngressCtl) Name() string {
	return "IngressCtl"
}

func (ing *IngressCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns","default")
	ob := ing.IngressMap.ListTest(ns)
	fmt.Println(ob)
	return gin.H{
		"code": 20000,
		"data": ob,
	}
}

func (ing *IngressCtl) Build(goft *goft.Goft) {
	goft.Handle("GET","/ingress",ing.ListAll)
}