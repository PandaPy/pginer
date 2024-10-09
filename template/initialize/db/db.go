package db

import (
	"log"

	"github.com/PandaPy/pginer/template/initialize/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBs map[string]*gorm.DB

// 初始化所有 MySQL 数据库连接
func InitDB() {
	DBs = make(map[string]*gorm.DB)
	for name, dbConfig := range config.AppConfig.Databases {

		dsn := dbConfig.Dsn()

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			zap.L().Fatal("failed to initialize database", zap.Error(err))
		}
		DBs[name] = db
	}
}

// 获取指定的 MySQL 数据库连接，默认返回 default 数据库
func DB(name ...string) *gorm.DB {
	if len(name) == 0 {
		name = append(name, "default")
	}
	db, exists := DBs[name[0]]
	if !exists {
		log.Fatalf("数据库 %v 不存在", name[0])
	}
	return db
}
