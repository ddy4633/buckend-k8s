package controller

import (
	"k8s-web/src/servers"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type PodCtl struct {
	PodServices *servers.PodService `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (p *PodCtl) Name() string {
	return "PodCtl"
}

func (p *PodCtl) GetALL(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": p.PodServices.ListPod(ns),
	}
}

func (p *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", p.GetALL)
}
