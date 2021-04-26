package controller

import (
	"k8s-web/src/wscore"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type WsCli struct {
}

func NewWsCli() *WsCli {
	return &WsCli{}
}

func (this *WsCli) Connect(c *gin.Context) (v goft.Void) {
	// 协议升级
	conn, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	} else {
		// 存储websocket对象
		wscore.ClientMap.Store(conn)
		return
	}
}

func (this *WsCli) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ws", this.Connect)
}

func (this *WsCli) Name() string {
	return "WsCli"
}
