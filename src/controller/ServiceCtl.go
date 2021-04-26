package controller

import (
	"k8s-web/src/servers"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
)

type ServiceCtl struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
	SVserver  *servers.SVservice    `inject:"-"`
}

func NewServiceCtl() *ServiceCtl {
	return &ServiceCtl{}
}

// 获取Services列表
func (s *ServiceCtl) Getlist(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": s.SVserver.ListAll(ns),
	}
}

func (s *ServiceCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/service", s.Getlist)
}

func (s *ServiceCtl) Name() string {
	return "ServiceCtl"
}
