package model

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"server/global"
	"server/internal/mysql"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddRecord(context *gin.Context) {
	var err error
	a := &mysql.RecordRow{}

	a.ModelID, err = strconv.Atoi(context.PostForm("ModelID"))
	if err != nil {
		global.Logger.Error("model id is not an int", err)
		panic(err)
	}

	eID, err := global.DB.SearchModelID(a.ModelID)
	if err != nil {
		panic(err)
	}
	if eID == "" {
		context.JSON(http.StatusOK, "associated employee id not found")
		eID = "unknown"
		//return
	}

	a.EmployeeID = eID

	t, err := time.Parse(time.RFC3339, context.PostForm("ScanTime"))
	if err != nil {
		global.Logger.Error("time format is wrong", err)
		panic(err)
	}
	a.ScanTime = t

	a.CameraID, err = strconv.Atoi(context.PostForm("CameraID"))
	if err != nil {
		global.Logger.Error("camera id is not an int", err)
		panic(err)
	}

	global.Logger.Info("added", a)

	fmt.Println(a)

	err = global.DB.Insert(a)
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, *a)

}

func AddPerson(context *gin.Context) {
	var err error

	e := &mysql.EmployeeRow{}
	r := &mysql.RecordRow{}

	eID := context.PostForm("EmployeeID")

	name, err := global.DB.SearchEmployeeID(eID)
	if err == nil {
		global.Logger.Error("search employee id failed", err)
		panic(err)
	}
	if name != "" {
		context.JSON(http.StatusOK, "person already exists")
		return
	}

	r.EmployeeID = eID
	e.EmployeeID = eID

	mID, err := strconv.Atoi(context.PostForm("ModelID"))
	if err != nil {
		global.Logger.Error("model id is not an int", err)
		panic(err)
	}

	name, err = global.DB.SearchModelID(mID)
	if err == nil {
		panic(err)
	}
	if name != "" {
		context.JSON(http.StatusOK, "person already exists")
		return
	}

	r.ModelID = mID
	e.ModelID = mID

	r.CameraID, err = strconv.Atoi(context.PostForm("CameraID"))
	if err != nil {
		global.Logger.Error("camera id is not an int", err)
		panic(err)
	}

	t, err := time.Parse(time.RFC3339, context.PostForm("ScanTime"))
	if err != nil {
		global.Logger.Error("time format is wrong", err)
		panic(err)
	}
	r.ScanTime = t


	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err = tx.Table("employee_info").Create(e).Error; err != nil {
			// return any error will rollback
			global.Logger.Error("create employee err", err)
			return err
		}

		if err = tx.Table("records").Create(r).Error; err != nil {
			global.Logger.Error("create record err", err)
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		panic(err)
	}

	//tx := global.DB.Begin()
	//
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//if err = tx.Error; err != nil {
	//	panic(err)
	//}
	//
	//if err = tx.Table("employee_info").Create(e).Error; err != nil {
	//	tx.Rollback()
	//	panic(err)
	//}
	//
	//if err = tx.Table("records").Create(r).Error; err != nil {
	//	tx.Rollback()
	//	panic(err)
	//}
	//
	//if err = tx.Commit().Error; err != nil {
	//	panic(err)
	//}

	context.JSON(http.StatusOK, *r)
}
