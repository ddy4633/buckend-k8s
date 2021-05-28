package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/deprecated/scheme"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"os"
)

func main() {
	config,_ := clientcmd.BuildConfigFromFlags("","/root/.kube/config")
	//config.Insecure = true
	resrclient,err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	options := &v1.PodExecOptions{
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
		Container: "centos-tools",
		Command:   []string{"sh"},
	}
	req := resrclient.CoreV1().RESTClient().Post().Resource("pods").
		Namespace("default").Name("centos-tools").
		SubResource("exec").VersionedParams(
			options,
			scheme.ParameterCodec)
	fmt.Println(req.URL())
	exec,err := remotecommand.NewSPDYExecutor(config,"POST",req.URL())
	if err !=nil{
		panic(err)
	}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:             os.Stdin,
		Stdout:            os.Stdout,
		Stderr:            os.Stderr,
		Tty:               true,
	})
	log.Println(err)

}