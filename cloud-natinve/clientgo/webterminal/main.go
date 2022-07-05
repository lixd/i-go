package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/klog/v2"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow connections from any Origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

/*
1. controller manager 传到 handler（请求里需要用到 client）
2. client 改为同时存储 client 和 config（exec 接口需要 config）
3. web terminal 逻辑
4. install 时增加 kubectl pod 步骤，deployment方式并限制cpu、memory
*/

/*
通过 apiserverr 提供的 exec 接口实现的 web terminal.
大致原理：https://github.com/lixd/daily-notes/blob/00a44c67abcfff9665f001da565bbdffcf50de37/CloudNative/Kubernetes/note/kubectl-console.md
具体实现参考：https://github.com/kubernetes/dashboard/blob/master/src/app/backend/handler/terminal.go
*/
const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// ctrl+d to close terminal
	endOfTransmission = "\u0004"
)

func main() {
	container := restful.NewContainer()
	err := AddToContainer(container)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	server := &http.Server{
		Addr: ":9090",
	}
	server.Handler = container
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func AddToContainer(c *restful.Container) error {
	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("/exec").
		To(handleTerminalSession).
		Param(webservice.PathParameter("namespace", "namespace of which the pod located in")).
		Param(webservice.PathParameter("pod", "name of the pod")).
		Doc("create terminal session").
		Metadata("123", nil).
		Writes(nil))

	c.Add(webservice)
	return nil
}

func handleTerminalSession(request *restful.Request, response *restful.Response) {
	conn, err := upgrader.Upgrade(response.ResponseWriter, request.Request, nil)
	if err != nil {
		klog.Warning(err)
		return
	}

	HandleSession(conn)
}

func HandleSession(conn *websocket.Conn) {
	// 1）加载配置文件
	// config, err := clientcmd.BuildConfigFromFlags("", clientgo.KubeConfig)
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err.Error())
	}

	// 2）实例化clientset对象
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	session := &TerminalSession{conn: conn, sizeChan: make(chan remotecommand.TerminalSize)}

	t := NewTerminaler(clientSet, config)
	err = t.startProcess("kube-system", "uk8s-kubectl-7bdd494f7-c8wp9", "uk8s-kubectl",
		[]string{"sh"}, session)
	if err != nil {
		session.Close(2, err.Error())
		return
	}

	session.Close(1, "Process exited")
}

type terminaler struct {
	client kubernetes.Interface
	config *rest.Config
}

func NewTerminaler(client kubernetes.Interface, config *rest.Config) *terminaler {
	return &terminaler{client: client, config: config}
}

// PtyHandler is what remotecommand expects from a pty
type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

// TerminalSession implements PtyHandler (using a SockJS connection)
type TerminalSession struct {
	conn     *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
}

func (t *terminaler) startProcess(namespace, podName, containerName string, cmd []string, ptyHandler PtyHandler) error {
	fmt.Printf("ns:%s pod:%s container:%s cmd:%v", namespace, podName, containerName, cmd)
	req := t.client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec")
	req.VersionedParams(&corev1.PodExecOptions{
		Container: containerName,
		Command:   cmd,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	fmt.Println("req.URL:", req.URL())
	exec, err := remotecommand.NewSPDYExecutor(t.config, "POST", req.URL())
	if err != nil {
		return errors.WithMessage(err, "NewSPDYExecutor")
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:             ptyHandler,
		Stdout:            ptyHandler,
		Stderr:            ptyHandler,
		TerminalSizeQueue: ptyHandler,
		Tty:               true,
	})
	if err != nil {
		return errors.WithMessage(err, "Stream")
	}

	return nil
}

type TerminalMessage struct {
	Op, Data   string
	Rows, Cols uint16
}

// Next handles pty->process resize events
// Called in a loop from remotecommand as long as the process is running
func (t TerminalSession) Next() *remotecommand.TerminalSize {
	size := <-t.sizeChan
	if size.Height == 0 && size.Width == 0 {
		return nil
	}
	return &size
}

// Read handles pty->process messages (stdin, resize)
// Called in a loop from remotecommand as long as the process is running
func (t TerminalSession) Read(p []byte) (int, error) {
	// 失败时发送 ctrl+c 命令关闭 remote terminal
	var msg TerminalMessage
	err := t.conn.ReadJSON(&msg)
	if err != nil {
		fmt.Println("readJSON err:", err)
		return copy(p, endOfTransmission), err
	}
	fmt.Printf("read data: %+v", msg)

	switch msg.Op {
	case "stdin":
		return copy(p, msg.Data), nil
	case "resize":
		t.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	default:
		return copy(p, endOfTransmission), fmt.Errorf("unknown message type '%s'", msg.Op)
	}
}

// Write handles process->pty stdout
// Called from remotecommand whenever there is any output
func (t TerminalSession) Write(p []byte) (int, error) {
	msg, err := json.Marshal(TerminalMessage{
		Op:   "stdout",
		Data: string(p),
	})
	if err != nil {
		return 0, err
	}
	fmt.Println("write data:", string(p))

	t.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err = t.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		fmt.Print("write err:", err)
		return 0, err
	}
	return len(p), nil
}

// Toast can be used to send the user any OOB messages
// hterm puts these in the center of the terminal
func (t TerminalSession) Toast(p string) error {
	msg, err := json.Marshal(TerminalMessage{
		Op:   "toast",
		Data: p,
	})
	if err != nil {
		return err
	}
	t.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err = t.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return err
	}
	return nil
}

func (t TerminalSession) Close(status uint32, reason string) {
	klog.Warning(status, reason)
	close(t.sizeChan)
	t.conn.Close()
}
