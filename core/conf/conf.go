// conf 加载配置文件
package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

// Init
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

// 配置文件初始化func
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("../conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("json") //  设置配置文件格式为json
	viper.AutomaticEnv()        // 读取匹配的环境变量

	viper.SetEnvPrefix("TESTSERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		log.Printf("config error : %s", err.Error())
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
