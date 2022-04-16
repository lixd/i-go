package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	pods()
	deployment()

	createPod()
	deletePod()
}

func pods() {
	// 1) 加载配置文件
	// kubeconfig := "~/.kube/config"
	kubeconfig := "D:\\Home\\17x\\Projects\\i-go\\cloud-natinve\\kubeconfig"
	// 从指定的URL加载kubeconfig配置文件
	// 这里直接从本机加载，因此第一个参数为空字符串
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil { // kubeconfig加载失败就直接退出了
		panic(err.Error())
	}
	// 2）构建client
	// 参考path : /api/v1/namespaces/{namespace}/pods
	config.APIPath = "api"
	// pod的group是空字符串
	config.GroupVersion = &corev1.SchemeGroupVersion
	// 指定序列化工具
	config.NegotiatedSerializer = scheme.Codecs

	// 根据配置信息构建restClient实例
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}
	// 3）组装请求并接收结果
	// 保存pod结果的数据结构实例
	result := &corev1.PodList{}
	// 参考path : /api/v1/namespaces/{namespace}/pods
	// 设置请求参数，然后发起请求
	err = restClient.Get(). // GET请求
				Namespace("kube-system").                                                //  指定namespace，
				Resource("pods").                                                        // 查找多个pod
				VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec). // 指定大小限制和序列化工具
				Do(context.TODO()).                                                      // 请求
				Into(result)                                                             // 结果存入result
	if err != nil {
		panic(err.Error())
	}
	// 4）处理结果--这里就打印结果到控制台
	// 表头
	fmt.Printf("namespace\t status\t\t name\n")
	// 每个pod都打印namespace、status.Phase、name三个字段
	for _, d := range result.Items {
		fmt.Printf("%v\t %v\t %v\n",
			d.Namespace,
			d.Status.Phase,
			d.Name)
	}
}

func deployment() {
	// 1) 加载配置文件
	// kubeconfig := "~/.kube/config"
	kubeconfig := "D:\\Home\\17x\\Projects\\i-go\\cloud-natinve\\kubeconfig"
	// 从指定的URL加载kubeconfig配置文件
	// 这里直接从本机加载，因此第一个参数为空字符串
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil { // kubeconfig加载失败就直接退出了
		panic(err.Error())
	}
	// 2）构建client
	// 参考path : /apis/apps/v1/namespaces/{namespace}/deployments
	config.APIPath = "apis"
	// deployment的group是apps
	config.GroupVersion = &appsv1.SchemeGroupVersion
	// 指定序列化工具
	config.NegotiatedSerializer = scheme.Codecs

	// 根据配置信息构建restClient实例
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}
	// 3）组装请求并接收结果
	// 保存pod结果的数据结构实例
	result := &appsv1.DeploymentList{}
	// 参考path : /apis/apps/v1/namespaces/{namespace}/deployments
	// 设置请求参数，然后发起请求
	err = restClient.Get(). // GET请求
				Namespace("kube-system").                                                //  指定namespace，
				Resource("deployments").                                                 // 查找多个pod
				VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec). // 指定大小限制和序列化工具
				Do(context.TODO()).                                                      // 请求
				Into(result)                                                             // 结果存入result
	if err != nil {
		panic(err.Error())
	}
	// 4）处理结果--这里就打印结果到控制台
	// 表头
	fmt.Printf("namespace\t replicas\t available\t name\n")
	for _, d := range result.Items {
		fmt.Printf("%v\t %v\t\t %v\t\t %v\n",
			d.Namespace,
			d.Status.Replicas,
			d.Status.AvailableReplicas,
			d.Name)
	}
}

func createPod() {
	kubeconfig := "D:\\Home\\17x\\Projects\\i-go\\cloud-natinve\\kubeconfig"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// 2）构建client
	// 参考path : POST /api/v1/namespaces/{namespace}/pods
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// 根据配置信息构建restClient实例
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}
	result := &corev1.Pod{}
	// 构建一个pod对象，这个还是挺麻烦的，建议对照着yaml来写
	/*
		比如这是一个简单的nginx pod yaml内容，照着转换成 corev1.pod 结构即可
		字段上有// +optional注释的则可以不填，其他的必须都填
		---
		apiVersion: v1
		kind: Pod
		metadata:
		  name: nginx
		  labels:
		    app: nginx
		spec:
		  containers:
		  - name: nginx
		    image: nginx:1.20.0
		    ports:
		    - containerPort: 80
	*/
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
	code := new(int)
	err = restClient.Post().
		Namespace("default"). // 这个namespace需要和pod里指定的相匹配，然后pod里的namespace未指定在会用这里的来填充
		Resource("pods").
		VersionedParams(&metav1.CreateOptions{}, scheme.ParameterCodec).
		Body(pod).
		Do(context.Background()).
		StatusCode(code).
		Into(result)
	if err != nil {
		log.Println("post err: ", err)
		return
	}
	// 根据返回状态码不同，代码不同含义
	switch *code {
	case 200:
		fmt.Println("OK")
	case 201:
		fmt.Println("Created")
	case 202:
		fmt.Println("Accepted")
	}
	fmt.Println(result)
}

func deletePod() {
	kubeconfig := "D:\\Home\\17x\\Projects\\i-go\\cloud-natinve\\kubeconfig"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// 2）构建client
	// 参考path : POST /api/v1/namespaces/{namespace}/pods
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// 根据配置信息构建restClient实例
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}
	err = restClient.Delete().
		Namespace("default").
		Resource("pods").
		Name("my-pod").
		Body(&metav1.DeleteOptions{}).
		Do(context.Background()).
		Error()
	if err != nil {
		log.Println("delete pod err: ", err)
	}
}

func updatePod() {
	kubeconfig := "D:\\Home\\17x\\Projects\\i-go\\cloud-natinve\\kubeconfig"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// 2）构建client
	// 参考path : POST /api/v1/namespaces/{namespace}/pods
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// 根据配置信息构建restClient实例
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}
	result := &corev1.Pod{}
	// 构建一个pod对象，这个还是挺麻烦的，建议对照着yaml来写
	pod := &corev1.Pod{
		// TypeMeta: metav1.TypeMeta{
		// 	Kind:       "pod",
		// 	APIVersion: "v1",
		// },
		// ObjectMeta: metav1.ObjectMeta{
		// 	Name: "my-pod",
		// 	Namespace: "default",
		// 	Labels:    map[string]string{"app": "nginx"},
		// },

		Spec: corev1.PodSpec{
			// InitContainers: []corev1.Container{},
			Containers: []corev1.Container{{
				Name:  "nginx",
				Image: "nginx:1.21", // 这里把镜像版本改一下
				// Ports: []corev1.ContainerPort{{
				// 	ContainerPort: 80,
				// }},
			}},
		},
		Status: corev1.PodStatus{},
	}
	code := new(int)
	// update 和 create 比较类似
	// POST方法改成PUT
	// 然后需要指定PodName
	err = restClient.Put().
		Namespace("default").
		Resource("pods").
		Name("my-pod").
		VersionedParams(&metav1.UpdateOptions{}, scheme.ParameterCodec).
		Body(pod).
		Do(context.Background()).
		StatusCode(code).
		Into(result)
	if err != nil {
		print(err.Error())
		// log.Println("put err: ", err)
		return
	}
	// 根据返回状态码不同，代码不同含义
	switch *code {
	case 200:
		fmt.Println("OK")
	case 201:
		fmt.Println("Created")
	default:
		fmt.Println("unknown code:", *code)
	}
	fmt.Println(result)
}
