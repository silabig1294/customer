package database

import (
	"context"
	"fmt"
	"time"
	"github.com/silabig1294/customer/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql,_ := fc()
	fmt.Printf("%v\n===========================\n",sql)


}

var DB *gorm.DB

func Connect() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/user?parseTime=true"
	dial := mysql.Open(dsn)
	connection , err := gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic("could not connect to the database")
	}
    
	DB = connection
	 
	connection.AutoMigrate(&models.User{})

}
