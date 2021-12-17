package main

import "fmt"

/*
https://avinetworks.com/glossary/subnet-mask/

IP 地址 = 网络地址 + 主机地址
子网掩码将IP地址拆分为主机地址和网络地址，从而定义IP地址的哪一部分属于设备，哪一部分属于网络。
子网掩码是通过将主机位设置为全 0 并将网络位设置为全 1 创建的 32 位数字。通过这种方式，子网掩码将 IP 地址分为网络地址和主机地址。
“255”地址总是分配给广播地址，“0”地址总是分配给网络地址。两者都不能分配给主机，因为它们是为这些特殊目的保留的。
一个网络又可以分为多个子网
IP地址 = 网络号 + 子网号 + 主机号
*/

// 使用二进制位来表示多种状态

const (
	System = 1
	Baidu  = System << 1
	Ali    = System << 2
)

// Check 存储目前的权限状态
type Check struct {
	Flag int
}

// SetStatus 设置状态
func (c *Check) SetStatus(status int) {
	c.Flag = status
}

// AddStatus 添加一种或多种状态
func (c *Check) AddStatus(status int) {
	c.Flag |= status
}

// DeleteStatus 删除一种或者多种状态
func (c *Check) DeleteStatus(status int) {
	/**
	go 不支持取反符号~
	c.Flag &= ~status
	取反 ^status
	*/
	c.Flag &= ^status
}

// HasStatus 是否具有某些状态
func (c *Check) HasStatus(status int) bool {
	return (c.Flag & status) == status
}

// NotHasStatus 是否不具有某些状态
func (c *Check) NotHasStatus(status int) bool {
	return (c.Flag & status) == 0
}

// OnlyHas 是否仅仅具有某些状态
func (c *Check) OnlyHas(status int) bool {
	return c.Flag == status
}

func main() {
	c := Check{Flag: 0}
	c.SetStatus(3) // 011
	fmt.Printf("3=>011 --> has system %t, and baidu %t, and ali %t\n", c.HasStatus(System), c.HasStatus(Baidu), c.HasStatus(Ali))
	c.AddStatus(Ali)
	fmt.Printf("3=>011 add ali --> has system %t, and baidu %t, and ali %t\n", c.HasStatus(System), c.HasStatus(Baidu), c.HasStatus(Ali))
	c.DeleteStatus(Baidu)
	fmt.Printf("3=>011 add ali and delete baidu --> has system %t, and baidu %t, and ali %t\n", c.HasStatus(System), c.HasStatus(Baidu), c.HasStatus(Ali))
	fmt.Printf("3=>011 add ali and delete baidu --> not has system %t, and not has baidu %t, and not has ali %t\n", c.NotHasStatus(System), c.NotHasStatus(Baidu), c.NotHasStatus(Ali))
	fmt.Printf("3=>011 add ali and delete baidu --> only has system %t\n", c.OnlyHas(System))
	c.DeleteStatus(Ali)
	fmt.Printf("3=>011 add ali and delete baidu and delete ali --> only has system %t\n", c.OnlyHas(System))
}
