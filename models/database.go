package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/web_qctx?charset=utf8mb4"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
