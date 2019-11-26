package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strconv"
)

func main() {
	// MsgPublish()
	var a = 3
	a -= 1
	fmt.Println(a)
}

// InitAndPublish 建立mqtt客户端，连接，并推送系统/私信消息（使用用户密码）
// param：topic  主题(一般为用户ID)
// param：msg  推送的消息
// param：quiesce  等待退出时长
// return: bool 是否成功
func MsgPublish(topic []int, msg string, uiesce uint) error {
	var err error
	opt := mqtt.NewClientOptions().AddBroker("ws://mqtt.cfwd.com:80").SetClientID("t1").SetUsername("admin").SetPassword("123456")
	client := mqtt.NewClient(opt)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("连接问题[%v]", token.Error())
		return token.Error()
	}
	for i := 0; i < len(topic); i++ {
		if token := client.Publish(strconv.Itoa(topic[i]), 0, false, msg); token.Wait() && token.Error() != nil {
			err = token.Error()
			continue
		}
	}
	client.Disconnect(uiesce)
	return err
}
