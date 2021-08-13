package router

import (
	"server/global"
	"server/model"
)

func SetUpRouter() {
	global.Router.POST("/add", model.AddPerson)
	global.Router.POST("/scan", model.AddRecord)
}
