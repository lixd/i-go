// Package conf 配置文件
package conf

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"i-go/core/db/elasticsearch"
	"i-go/data"
	"i-go/data/conf"
	"i-go/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Loads 手动调用加载配置文件
func Loads(files []string) error {
	for _, file := range files {
		_ = Load(file)
	}
	return nil
}

// Load 手动调用加载配置文件
func Load(file string) error {
	if runtime.GOOS == "windows" {
		file = data.Path(file)
	}
	// 初始化配置文件
	if err := initConfig(file); err != nil {
		return err
	}
	// 监控配置文件变化并热加载程序
	watchConfig()
	return nil
}

// initConfig 配置文件初始化
func initConfig(file string) error {
	viper.SetConfigFile(file)  // 指定配置文件
	viper.SetConfigType("yml") //  设置配置文件格式为yml
	viper.AutomaticEnv()       // 读取匹配的环境变量

	viper.SetEnvPrefix("TEST_") // 读取环境变量的前缀为TEST_
	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}
	return nil
}

// watchConfig 监控配置文件变化并热加载程序
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)

		prefix := utils.GetFilePrefix(e.Name)
		switch prefix {
		case conf.Elasticsearch:
			fmt.Println("elasticsearch conf changed!")
			// 配置文件更新后再次初始化
			elasticsearch.Init()
		case conf.MongoDB:
			fmt.Println("mongo conf changed!")
		case conf.Redis:
			fmt.Println("redis conf changed!")
		case conf.Basic:
			fmt.Println("basic conf changed!")
		default:
			fmt.Println("conf changed!")
		}
	})
}
