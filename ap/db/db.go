package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Open() {
	connect := "eapp:password@tcp(db:3306)/esample"
	db, err = gorm.Open("mysql", connect)
	// defer db.Close()
	if err != nil {
		panic(err.Error())
	}
}

func Conn() *gorm.DB {
	return db
}
