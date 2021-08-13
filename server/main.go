package main

import (
	"server/global"
	"server/initialize"
	"server/router"
)

func main() {
	initialize.ViperRead()
	initialize.InitBeegoLogger()
	global.DB = initialize.GormMySQL()

	router.SetUpRouter()

	global.Logger.Info("service start!")
	global.Logger.Error("service start![error]")
	global.Router.Run(":" + global.Config.Sys.Port)
}
