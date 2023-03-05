package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect()  {
	dsn := "root:admin@tcp(127.0.0.1:3306)/user?parseTime=true"
	dial := mysql.Open(dsn)
	_,err := gorm.Open(dial,&gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

}