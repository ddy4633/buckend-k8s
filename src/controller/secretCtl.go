package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-web/src/servers"
)

type SecretCtl struct {
	SecretService *servers.SecretService `inject:"-"`
}

// 获取全部secret对象
func (se *SecretCtl) GetAll(c *gin.Context) goft.Json {
	return  gin.H{
		"code": 20000,
		"data": se.SecretService.Getall(c.DefaultQuery("ns","default")),
	}
}

// 删除指定的secret对象
func (se *SecretCtl) DeleteSecret(c *gin.Context) goft.Json  {
	return gin.H{
		"code": 20000,
		"data": se.SecretService.DeleteSecret(
			c.DefaultQuery("ns","default"),
			c.DefaultQuery("name"," ")),
	}

}

// 创建secret对象
func (se *SecretCtl) CreateSecret(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": 0,
	}
}

func NewSecretCtl() *SecretCtl {
	return &SecretCtl{}
}

func (se *SecretCtl) Name() string {
	return "SecretCtl"
}

func (se *SecretCtl) Build(goft *goft.Goft) {
	goft.Handle("GET","/secret",se.GetAll)
	goft.Handle("DELETE","/secret",se.DeleteSecret)
	goft.Handle("POST","/secret",se.CreateSecret)
}