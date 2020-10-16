package mysqlConnector

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@(129.211.93.165:3306)/imilateqq?charset=utf8")
	if err != nil {
		fmt.Println(err)
		fmt.Println(db)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.SingularTable(true)
}

func GetDBConn() *gorm.DB {
	return db
}
