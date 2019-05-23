package main

import "fmt"

// 声明/定义一个接口
type Usb interface {
	// 声明两个未实现的方法
	Start()
	Stop()
}

// 手机
type Phone struct {
}

// 让手机实现usb接口中的方法
func (phone Phone) Start() {
	fmt.Println("Phone Start..")
}
func (phone Phone) Stop() {
	fmt.Println("Phone Stop..")
}
func (phone *Phone) Call() {
	fmt.Println("打电话。。。")
}

// 相机
type Camera struct {
}

// 让相机实现usb接口中的方法
func (camera Camera) Start() {
	fmt.Println("Camera Start..")
}
func (camera Camera) Stop() {
	fmt.Println("Camera Stop..")
}

// 让手机实现usb接口中的方法
type Computer struct {
}

// 编写一个方法Working 方法接收Usb接口类型
// 只要是实现可 Usb 接口(所谓Usb接口就是指实现了Usb接口声明的所有方法)
func (computer Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}
func main() {
	// 创建变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	// 测试 关键点
	computer.Working(phone)
	fmt.Println("----------------")
	computer.Working(camera)

	var x float32
	var y interface{}
	y = x // 空接口可以接收任何类型
	// x = y           // 报错 x是接口类型无法直接赋值给t 使用类型断言
	// 类型断言(带检测的)
	t1, ok := y.(float64) // 转成float32
	if ok {
		fmt.Println(t1)
	} else {
		fmt.Println("类型断言失败")
	}
	TypeJudg("xxx")
}

func TypeJudg(items ...interface{}) {
	for i, value := range items {
		//  value.(type) 其中 type 是固定写法
		switch value.(type) {
		case bool:
			fmt.Printf("参数 %d 值是 %v 类型是 bool", i, value)
		case string:
			fmt.Printf("参数 %d 值是 %v 类型是 string", i, value)
		}

	}
}
