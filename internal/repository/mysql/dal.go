package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"niceBackend/config"
	"niceBackend/pkg/util"
	"sync"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DB = ConnectDB()
	})
}

func ConnectDB() (conn *gorm.DB) {
	var err error
	confFile := util.GetConfigIniPath()
	errI := config.Init(confFile)
	if errI != nil {
		panic(fmt.Errorf("init config is error"))
	}
	conf := config.GetConf()
	m := conf.Mysql.Write
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	conn, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
