package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"io"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"net/http"
	"strconv"
	"time"
)

type PodLogCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
}

func NewPodLogCtl() *PodLogCtl {
	return &PodLogCtl{}
}

func (plog *PodLogCtl) GetLogs(c *gin.Context) (v goft.Void) {
	// 分别获取namespace，pod，containter
	ns := c.DefaultQuery("ns","default")
	podname := c.DefaultQuery("pname","")
	ctname := c.DefaultQuery("cname","")
	tail,_ := strconv.ParseInt(c.DefaultQuery("tail","500"),10,64)
	// 日志的请求参数(指定的container+实时读取+默认500行)
	opt := &coreV1.PodLogOptions{Container: ctname,TailLines: &tail,Follow: true}
	// 获取容器的日志信息
	reqpodlogs := plog.Client.CoreV1().Pods(ns).GetLogs(podname,opt)
	log.Println(ns,podname,ctname,opt)
	// 设置单个超时的时间为10分钟
	cc,_ := context.WithTimeout(c,5*time.Minute)
	// 流式读取日志信息
	ioreader,err := reqpodlogs.Stream(cc)
	// 这里如果pod资源不存在会抛出指针错误待修复
	if err !=nil {
		log.Printf("%s/%s -> %s Read logs Error: %s",ns,podname,ctname,err)
		return
	}
	defer ioreader.Close()
	for {
		buf := make([]byte,1024)
		n,err := ioreader.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		if n>0 {
			//fmt.Println(string(buf[0:n]))
			c.Writer.Write([]byte(string(buf[0:n])))
			c.Writer.(http.Flusher).Flush()
		}
	}
	return
}

func (*PodLogCtl) Name() string{
	return "PodLogCtl"
}

func (plog *PodLogCtl) Build(goft *goft.Goft){
	goft.Handle("GET","/pods/logs",plog.GetLogs)
}

