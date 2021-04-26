package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CrosMiddleware struct {
}

func NewCrosMiddleware() *CrosMiddleware {
	return &CrosMiddleware{}
}

func (*CrosMiddleware) OnRequest(ctx *gin.Context) error {
	method := ctx.Request.Method
	// 加入规范头文件
	if method != "" {
		ctx.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
	}
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	return nil
}

func (*CrosMiddleware) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}
