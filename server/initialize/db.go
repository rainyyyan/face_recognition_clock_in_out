package initialize

import (
	"server/global"
	in "server/internal/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMySQL() *in.FRDB {
	dsn := global.Config.DB.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	a := &in.FRDB{db}
	global.Logger.Info("connected to database")
	return a
}
