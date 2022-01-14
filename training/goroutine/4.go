package main

import _ "go.uber.org/automaxprocs"

// github.com/uber-go/automaxprocs 容器环境 自动配置 maxprocs
func main() {
	// 自动获取cgroup中的数据来设置maxprocess
	// 具体为 cpu.cfs_quota_us / cpu.cfs_period_us
}
