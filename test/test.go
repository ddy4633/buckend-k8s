package main

import (
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/deprecated/scheme"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"os"
	"strings"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("","/root/.kube/config" )
	if err!=nil{
		log.Fatal(err)
	}
	//config.Insecure=true
	client,err:=kubernetes.NewForConfig(config)
	if err!=nil{
		log.Fatal(err)
	}

	r:=gin.New()
	r.POST("/", func(c *gin.Context) {
		body,_:=c.GetRawData() //获取body 原始内容
		cmd:=strings.Split(string(body)," ")
		err=HandleCommand(client,config,cmd).Stream(remotecommand.StreamOptions{
			Stdout:c.Writer,
			Stderr:os.Stderr,
			Tty:true,
		})
	})
	r.Run(":8080")
}


func HandleCommand(client *kubernetes.Clientset,config *rest.Config,command []string) remotecommand.Executor{
	option := &v1.PodExecOptions{
		Container:"centos-tools",
		Command: []string{"sh"},
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:true,
	}
	req:=client.CoreV1().RESTClient().Post().Resource("pods").
		Namespace("default").
		Name("centos-tools").
		SubResource("exec").
		Param("color","false").
		VersionedParams(
			option,
			scheme.ParameterCodec,
		)

	exec,err:=remotecommand.NewSPDYExecutor(config,"POST",
		req.URL())
	if err!=nil{
		panic(err)
	}
	return exec
}