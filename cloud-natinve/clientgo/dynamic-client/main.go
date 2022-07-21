package main

import (
	"context"
	"fmt"

	"i-go/cloud-natinve/clientgo"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1) 配置 kubeConfig
	config, err := clientcmd.BuildConfigFromFlags("", clientgo.KubeConfig)
	if err != nil {
		panic(err.Error())
	}
	// 2）初始化 dynamicClient 示例
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// 3）组装请求并执行
	// dynamicClient的唯一关联方法所需的入参为 GVR，因为需要用GVR来确认具体资源类型
	gvr := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}

	// 使用dynamicClient的查询列表方法，查询指定namespace下的所有pod，
	// 注意此方法返回的数据结构类型是UnstructuredList
	unstructObj, err := dynamicClient.
		Resource(gvr).
		Namespace("kube-system").
		List(context.TODO(), metav1.ListOptions{Limit: 100})

	if err != nil {
		panic(err.Error())
	}

	// 4）将 Unstructured 结构转成对应的资源类型
	podList := &apiv1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), podList)
	if err != nil {
		panic(err.Error())
	}
	// 5）处理结果
	// 表头
	fmt.Printf("namespace\t status\t\t name\n")

	// 每个pod都打印namespace、status.Phase、name三个字段
	for _, d := range podList.Items {
		fmt.Printf("%v\t %v\t %v\n",
			d.Namespace,
			d.Status.Phase,
			d.Name)
	}
}
