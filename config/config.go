/*
全局配置服务

BetaX Server Monitor
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const VERSION = "1.0.0"

func Init() {
	viper.SetConfigName("smc")
	viper.SetConfigType("ini")
	// viper.AddConfigPath("/etc/betax/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createDefault()
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println(err)
		}
	}
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
	viper.WriteConfig()
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func createDefault() {
	// 服务进程号
	viper.SetDefault("service.pid", "0")
	// 服务端口
	viper.SetDefault("service.port", "12780")
	// 报送服务故障时最大重启次数
	viper.SetDefault("service.max-retry", "3")
	// 服务上报速度(秒)
	viper.SetDefault("service.rate", "10")
	// 报送服务开关
	viper.SetDefault("service.delivery", "true")
	// 数据接口开关
	viper.SetDefault("service.api", "true")
	// 主控地址
	viper.SetDefault("master.addr", "localhost")
	// 主控端口
	viper.SetDefault("master.port", "12780")
	// 接入密钥
	viper.SetDefault("master.key", "")
	viper.SafeWriteConfig()
}
