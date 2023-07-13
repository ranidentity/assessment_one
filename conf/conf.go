package conf

import (
	"os"
	"singo/model"
	"singo/util"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/en.yaml"); err != nil {
		util.Log().Panic("Translation file loading error", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	// cache.Redis() redis not needed
}
