// 集群外认证并操作集群
package main

import (
    "context"
    "log"
    "time"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)

func main(){
    // 集群内pod会挂载认证信息
    config,err:=rest.InClusterConfig()
    if err != nil {
        log.Fatal(err) 
    }
    // 创建client
    clientset,err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatal(err)
    }
    // 每隔5秒使用client查询集群中所有的pod
    for {
        pods,err := clientset.CoreV1().Pods("default").List(context.TODO(),metav1.ListOptions{})
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("There are %d pods in the cluster\n",len(pods.Items))
        for i,pod := range pods.Items {
            log.Printf("%d -> %s/%s",i+1,pod.Namespace,pod.Name)
        }
        <-time.Tick(5*time.Second)
    }
}