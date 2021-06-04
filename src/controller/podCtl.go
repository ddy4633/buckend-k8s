package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-web/src/servers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodCtl struct {
	PodServices *servers.PodService `inject:"-"`
	Helper *servers.Helper `inject:"-"`
	ClientSet *kubernetes.Clientset `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (p *PodCtl) Name() string {
	return "PodCtl"
}

// 获取所有容器
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
			ipod).SetEXT(gin.H{"ReadNum":readycount,"Totle":totlecount}),
	}
}

// 获取容器组
func (p *PodCtl) Containers(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns","default")
	podname := c.DefaultQuery("name","")
	return gin.H{
		"code": 20000,
		"data": p.PodServices.GetContainers(ns,podname),
	}
}

// 删除指定的pod
func (p *PodCtl) DeletePod(c *gin.Context) goft.Json {
	ns := c.Query("ns")
	pod := c.Query("pod")
	err := p.ClientSet.CoreV1().Pods(ns).Delete(context.TODO(),pod,metav1.DeleteOptions{})
	return gin.H{
		"code": 20000,
		"data": err.Error(),
	}
}

func (p *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", p.GetALL)
	goft.Handle("GET", "/pods/containers", p.Containers)
	goft.Handle("DELETE","/pods/delete",p.DeletePod)
}
