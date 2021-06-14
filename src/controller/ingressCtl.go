package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-web/src/models"
	"k8s-web/src/servers"
)

type IngressCtl struct {
	IngressService *servers.IngressService `inject:"-"`
}

func NewIngressCtl() *IngressCtl {
	return &IngressCtl{}
}

func (*IngressCtl) Name() string {
	return "IngressCtl"
}

// 获取当前的ingress
func (ing *IngressCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns","default")
	return gin.H{
		"code": 20000,
		"data": ing.IngressService.GetALLIngress(ns),
	}
}

// 获取ingress的Annotation信息
func (ing *IngressCtl) GetAnnotations(c *gin.Context) goft.Json {
	//ns := c.DefaultQuery("ns","default")
	return gin.H{
		"code": 20000,
		"data": models.IngressAnnotationsitem,
	}
}

// 创建前端Post回来的Ingress
func (ing *IngressCtl) CreateIngress(c *gin.Context) goft.Json {
	ingressModles := &models.IngressCreate{}
	if err :=c.BindJSON(ingressModles);err != nil {
		fmt.Println(err)
		goft.Error(err)
	}
	return gin.H{
		"code": 20000,
		"data": ing.IngressService.CreateIngress(ingressModles),
	}
}


func (ing *IngressCtl) Build(goft *goft.Goft) {
	goft.Handle("GET","/ingress",ing.ListAll)
	goft.Handle("GET","/ingress/annotations",ing.GetAnnotations)
	goft.Handle("POST","/ingress",ing.CreateIngress)
}