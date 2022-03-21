package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

const MySQLDSN = "root:123456@(localhost:3306)/nice?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DB = ConnectDB()
	})
}

func ConnectDB() (conn *gorm.DB) {
	var err error
	conn, err = gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
