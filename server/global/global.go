package global

import (
	"github.com/gin-gonic/gin"
	"server/config"
	"server/internal/mysql"

	"github.com/beego/beego/v2/core/logs"
)

var (
	Config config.Server
	DB     *mysql.FRDB
	Logger *logs.BeeLogger
	Router = gin.Default()
)
