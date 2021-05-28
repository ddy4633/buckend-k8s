package additional

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/deprecated/scheme"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

// 处理Container的command执行
func HandlePodCommand(ns,pod,container string,client *kubernetes.Clientset,config *rest.Config,command []string) remotecommand.Executor {
	// 填充参数
	options := &coreV1.PodExecOptions{
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
		Container: container,
		Command:   command,
	}
	req := client.CoreV1().RESTClient().Post().Resource("pods").
		Namespace(ns).Name(pod).
		SubResource("exec").VersionedParams(
		options,
		scheme.ParameterCodec)
	fmt.Println(req.URL())
	exec,err := remotecommand.NewSPDYExecutor(config,"POST",req.URL())
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return exec
}
