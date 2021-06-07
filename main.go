package main

import (
	"k8s-web/src/config"
	controller "k8s-web/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	// 初始化脚手架
	server := goft.Ignite(cross()).Config(
		config.NewK8sHandle(),
		config.NewK8sConfig(),
		config.NewMaps(),
		config.NewServiceConfig(),
	).
		Mount("v1",
			controller.NewDeploymentCtl(),
			controller.NewPodCtl(),
			controller.NewUserCtl(),
			controller.NewWsCli(),
			controller.NewNsCtl(),
			controller.NewServiceCtl(),
			controller.NewPodLogCtl(),
			controller.NewIngressCtl(),
		).Attach(
	//middleware.NewCrosMiddleware()
	)
	//server.GET("/admin/*filepath", func(c *gin.Context) {
	//	// 转换成纯粹的gin的方式
	//	http.FileServer(FS(false)).ServeHTTP(c.Writer, c.Request)
	//})
	server.Launch()
}

// 解决跨域问题
func cross() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
