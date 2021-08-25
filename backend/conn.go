package backend

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/cjdjczym/TempBackend/model"
)


const (
	user   = "root"
	passwd = ""
	dbAddr = "127.0.0.1:3306"
	dbName = "test"
)

func connect() *gorm.DB {
	dsn := user + ":" + passwd + "@tcp(" + dbAddr + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("connect to mysql failed, err: " + err.Error())
		return nil
	}
	return db
}

func (u *model.UserDail)PostUserDaily() {
}
