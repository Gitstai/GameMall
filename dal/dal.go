package dal

import (
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	HOST     = "127.0.0.1:3306"
	DBName   = "gamemall"
)

var EduDB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, DBName)
	var err error
	EduDB, err = gorm.Open("mysql", dsn)
	if err != nil {
		logger.Error("InitDB err:", err)
		return err
	}
	EduDB.SingularTable(true)
	EduDB.LogMode(true)
	logger.Error("InitDB success")
	return nil
}
