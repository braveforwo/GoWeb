package connector

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@(xxx.xxx.xxx.xxx:3306)/imilateqq?charset=utf8")
	if err != nil {
		fmt.Println(err)
		fmt.Println(db)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.SingularTable(true)
}

func GetDBConn() *gorm.DB {
	return db
}
