package initialize

import (
	"fmt"
	"server/global"

	"github.com/beego/beego/v2/core/logs"
)

func InitBeegoLogger() {
	global.Logger = logs.NewLogger(64)
	global.Logger.SetLogger("console", "")
	beegoLogger(global.Config.Log.FilePath, global.Config.Log.Level)
}

func beegoLogger(logPath, level string) {
	configs := fmt.Sprintf(`{"filename":"%s","daily":true,"maxdays":7,"rotate":true}`, logPath)
	global.Logger.SetLogger("file", configs)
	switch level {
	case "Debug":
		global.Logger.SetLevel(logs.LevelDebug)
	case "Info":
		global.Logger.SetLevel(logs.LevelInfo)
	default:
		global.Logger.SetLevel(logs.LevelInfo)
	}

	global.Logger.EnableFuncCallDepth(true)
	global.Logger.SetLogFuncCallDepth(2)
}
