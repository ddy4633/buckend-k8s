package controller

import (
	"k8s-web/src/servers"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
)

type DeploymentCtl struct {
	K8sClient  *kubernetes.Clientset      `inject:"-"`
	DepService *servers.DeploymentService `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

// 获取deployment列表
func (this *DeploymentCtl) Getlist(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": this.DepService.ListAll(ns),
	}
}

func (this *DeploymentCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/deployment", this.Getlist)
}

func (this *DeploymentCtl) Name() string {
	return "DeploymentCtl"
}
