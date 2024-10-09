package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var AppConfig Config

// LoadConfig loads the configuration for the entire application
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		zap.L().Fatal("fatal error config file", zap.Error(err))
	}

	// 将配置内容解析到结构体
	if err := viper.Unmarshal(&AppConfig); err != nil {
		zap.L().Fatal("[unable to decode into struct]", zap.Error(err))
	}

	// 参数存在性校验
	validateConfig()
}

func validateConfig() {
	requiredKeys := []string{
		"LISTEN",
		"MODE",
		"ALLOWED_HOSTS",
		"SECRET_KEY",
		"DATABASES.default.NAME",
		"DATABASES.default.USER",
		"DATABASES.default.PASSWORD",
		"DATABASES.default.HOST",
		"DATABASES.default.PORT",
		"REDIS.HOST",
		"REDIS.USER",
		"REDIS.PASSWORD",
	}

	// 检查每一个必需的键是否在配置文件中设置
	for _, key := range requiredKeys {
		if !viper.IsSet(key) {
			zap.L().Fatal("config validation error", zap.Error(fmt.Errorf("missing required configuration key: %s", key)))
		}
	}
}
