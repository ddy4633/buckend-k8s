package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-web/src/servers"
)

type EndpointsCtl struct {
	EndpointsSV *servers.EndPointSV `inject:"-"`
}

func NewEndpointsCtl() *EndpointsCtl {
	return &EndpointsCtl{}
}

// 获取deployment列表
func (ep *EndpointsCtl) Getlist(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": ep.EndpointsSV.Getall(c.DefaultQuery("ns","default")),
	}
}

func (ep *EndpointsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/endpoints", ep.Getlist)
}

func (ep *EndpointsCtl) Name() string {
	return "EndPoints"
}

