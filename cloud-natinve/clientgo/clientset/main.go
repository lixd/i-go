package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"i-go/cloud-natinve/clientgo"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/pointer"
)

const (
	namespace      = "test-clientset"
	deploymentName = "client-test-deployment"
	serviceName    = "client-test-service"
)

var operate string

// go run main.go -operate create
// go run main.go -operate clean
func init() {
	// 获取用户输入的操作类型，默认是create，还可以输入clean，用于清理所有资源
	flag.StringVar(&operate, "operate", "create", "operate type : create or clean")
	flag.Parse()
	fmt.Printf("operation is %v\n", operate)
}

func main() {
	tomcat()
}

func tomcat() {
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

	// 3）根据参数执行对应操作
	switch operate {
	case "create":
		createNamespace(clientSet)
		createDeployment(clientSet)
		createService(clientSet)
	case "clean":
		clean(clientSet)
	default:
		fmt.Println("unknown operate")
	}
}

// 清理本次实战创建的所有资源
func clean(clientSet *kubernetes.Clientset) {
	// 删除service
	if err := clientSet.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{}); err != nil {
		log.Println("delete svc err:", err)
	}
	// 删除deployment
	if err := clientSet.AppsV1().Deployments(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{}); err != nil {
		log.Println("delete deployment err:", err)
	}
	// 删除namespace
	if err := clientSet.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{}); err != nil {
		log.Println("delete ns err:", err)
	}
}

// 新建namespace
func createNamespace(clientSet *kubernetes.Clientset) {
	namespaceClient := clientSet.CoreV1().Namespaces()
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	result, err := namespaceClient.Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Create namespace %s \n", result.GetName())
}

// 新建service
func createService(clientSet *kubernetes.Clientset) {

	// 得到service的客户端
	serviceClient := clientSet.CoreV1().Services(namespace)
	// 实例化一个数据结构
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{
				Name:     "http",
				Port:     8080,
				NodePort: 30080,
			},
			},
			Selector: map[string]string{
				"app": "tomcat",
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}

	result, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Create service %s \n", result.GetName())
}

// 新建deployment
func createDeployment(clientSet *kubernetes.Clientset) {
	// 得到deployment的客户端
	deploymentClient := clientSet.AppsV1().Deployments(namespace)

	// 实例化一个数据结构
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "tomcat",
				},
			},

			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "tomcat",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "tomcat",
							Image:           "tomcat:8.0.18-jre8",
							ImagePullPolicy: "IfNotPresent",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolSCTP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deploymentClient.Create(context.TODO(), deployment, metav1.CreateOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Create deployment %s \n", result.GetName())
}
