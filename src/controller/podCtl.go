package controller

import (
	"k8s-web/src/servers"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type PodCtl struct {
	PodServices *servers.PodService `inject:"-"`
	Helper *servers.Helper `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (p *PodCtl) Name() string {
	return "PodCtl"
}

func (p *PodCtl) GetALL(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	// 当前页面
	page := c.DefaultQuery("current","1")
	// 每一页的pod数量
	size := c.DefaultQuery("size","8")
	// 获取排序后的最终的Pod列表
	pods := p.PodServices.ListPod(ns)
	// 获取pod总数和就绪的总数
	readycount,totlecount,ipod :=p.PodServices.GetPodtotle(pods)
	return gin.H{
		"code": 20000,
		"data": p.Helper.PageResource(
			p.Helper.StrTOint(page,1),
			p.Helper.StrTOint(size,8),
			len(pods),ipod).SetEXT(gin.H{"ReadNum":readycount,"Totle":totlecount}),
	}
}

func (p *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", p.GetALL)
}
