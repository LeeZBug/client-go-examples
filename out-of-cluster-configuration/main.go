// 集群外认证并操作集群
package main

import (
    "context"
    "log"
    "path/filepath"
    "time"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/homedir"
)

func main() {
    // 获取用户目录
    homePath := homedir.HomeDir()
    if homePath == "" {
        log.Fatal("failed to get home directory!")
    }
    // 获取kubeconfig文件
    kubeconfig := filepath.Join(homePath, ".kube", "config")
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        log.Fatal(err)
    }
    // 创建client
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatal(err)
    }
    // 每隔5秒使用client查询集群中所有的pod
    for {
        pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("There are %d pods in the cluster\n", len(pods.Items))
        for i, pod := range pods.Items {
            log.Printf("%d -> %s/%s", i+1, pod.Namespace, pod.Name)
        }
        <-time.Tick(5 * time.Second)
    }
}
