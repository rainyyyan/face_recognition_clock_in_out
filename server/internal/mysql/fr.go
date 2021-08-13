package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

type FRDB struct {
	*gorm.DB
}

type EmployeeRow struct {
	EmployeeID 	string
	Name 		string
	ModelID     int
}

type RecordRow struct {
	EmployeeID string
	ModelID    int
	CameraID   int
	ScanTime   time.Time
}

func CreateDBSession(c *DBConfig) (*FRDB, error) {
	dsn := makeDSN(c)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	a := &FRDB{db}
	return a, nil
}

func makeDSN(c *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		c.Username, c.Password, c.Host, c.Port, c.Database)
}

func (RecordRow) TableName() string {
	return "records"
}

func (EmployeeRow) TableName() string {
	return "employee_info"
}

func (db *FRDB) Insert(r *RecordRow) error {
	result := db.Table("records").Create(r)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return fmt.Errorf("rows affected not 1")
	}
	return nil
}

func (db *FRDB) Add(e *EmployeeRow) error {
	result := db.Table("employee_info").Create(e)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return fmt.Errorf("rows affected not 1")
	}
	return nil
}

func (db *FRDB) SearchEmployeeID(eID string) (string, error) {
	var employee EmployeeRow
	result := db.Table("employee_info").Where("employee_id = ?", eID).Find(&employee)
	return employee.Name, result.Error
}

func (db *FRDB) SearchModelID(mID int) (string, error) {
	var employee EmployeeRow
	result := db.Table("employee_info").Where("model_id = ?", mID).Find(&employee)
	return employee.Name, result.Error
}