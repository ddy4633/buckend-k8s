package controller

import (
	"fmt"
	"k8s-web/src/additional"
	"k8s-web/src/wscore"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type WsCli struct {
	Client *kubernetes.Clientset `inject:"-"`
	Config *rest.Config `inject:"-"`
}

func NewWsCli() *WsCli {
	return &WsCli{}
}

// 常规资源的websocket连接
func (ws *WsCli) Connect(c *gin.Context) (v goft.Void) {
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

// 建立pod的连接
func (ws *WsCli) PodexecConnect(c *gin.Context) (v goft.Void) {
	ns := c.Query("ns")
	pod := c.Query("pod")
	container := c.Query("cname")
	terms := c.Query("terms")
	fmt.Println(terms)
	wsclient,err := wscore.Upgrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil {
		log.Println(err)
		return
	}
	shellClient := wscore.NewWebShellClient(wsclient)
	additional.HandlePodCommand(ns,pod,container,ws.Client,ws.Config,[]string{terms}).
		Stream(remotecommand.StreamOptions{
			Stdin:             shellClient,
			Stdout:            shellClient,
			Stderr:            shellClient,
			Tty:               true,
		})
	return
}

func (ws *WsCli) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ws", ws.Connect)
	goft.Handle("GET","/execws",ws.PodexecConnect)
}

func (ws *WsCli) Name() string {
	return "WsCli"
}
