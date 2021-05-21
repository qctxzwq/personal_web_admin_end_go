package models

import (
	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var Db *gorm.DB

func InitDb() (err error) {
	dsn := "root:123456@tcp(1.116.174.80:3306)/web_qctx?charset=utf8mb4"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

// 分页
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		r.ParseForm()
		var page, pageSize int
		if len(r.Form["page"]) == 0 {
			page = 1
		} else {
			page, _ = strconv.Atoi(string(r.Form["page"][0]))
		}

		if len(r.Form["psize"]) == 0 {
			pageSize = 100
		} else {
			pageSize, _ = strconv.Atoi(string(r.Form["psize"][0]))
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// 用户名
func FilterName(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		r.ParseForm()
		beego.Debug("进来了进来了", r.Form["name"])
		if len(r.Form["name"]) != 0 && len(r.Form["name"][0]) > 0 {
			return db.Where("name LIKE ?", "%"+r.Form["name"][0]+"%")
		}
		return db
	}
}

// 用户ID
func FilterId(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		r.ParseForm()
		beego.Debug("进来了进来了", r.Form["id"])
		if len(r.Form["id"]) != 0 && len(r.Form["id"][0]) > 0 {
			return db.Where("id LIKE ?", "%"+r.Form["id"][0]+"%")
		}
		return db
	}
}
