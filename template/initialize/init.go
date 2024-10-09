package initialize

import (
	"github.com/PandaPy/pginer/template/initialize/config"
	"github.com/PandaPy/pginer/template/initialize/db"
	"github.com/PandaPy/pginer/template/initialize/logger"
)

func Init() {
	// 初始化日志
	logger.InitLogger()

	// 初始化配置
	config.InitConfig()

	// 初始化数据库
	db.InitDB()
}
