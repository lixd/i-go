package main

import (
	"context"
	"fmt"
	"log"

	"i-go/cloud-natinve/clientgo"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	refactorPod()
}

func refactorPod() {
	// 1）加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", clientgo.KubeConfig)
	if err != nil {
		panic(err.Error())
	}

	// 2）实例化clientset对象
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// 	3) 具体逻辑
	listPod(clientSet)
	createPod(clientSet)
	deletePod(clientSet)
}

func listPod(clientSet *kubernetes.Clientset) {
	list, err := clientSet.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return
	}
	fmt.Printf("namespace\t status\t\t name\n")
	for _, d := range list.Items {
		fmt.Printf("%v\t %v\t %v\n",
			d.Namespace,
			d.Status.Phase,
			d.Name)
	}
}

func createPod(clientSet *kubernetes.Clientset) {
	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod",
			Namespace: "default",
			Labels:    map[string]string{"app": "nginx"},
		},
		Spec: corev1.PodSpec{
			InitContainers: []corev1.Container{},
			Containers: []corev1.Container{{
				Name:  "nginx",
				Image: "nginx:1.20",
				Ports: []corev1.ContainerPort{{
					ContainerPort: 80,
				}},
			}},
		},
		Status: corev1.PodStatus{},
	}
	result, err := clientSet.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		log.Println("create pod err:", err.Error())
		return
	}
	fmt.Println("create pod success:", result)
}

func deletePod(clientSet *kubernetes.Clientset) {
	err := clientSet.CoreV1().Pods("default").Delete(context.TODO(), "my-pod", metav1.DeleteOptions{})
	if err != nil {
		log.Println("delete pod err:", err.Error())
		return
	}
}
