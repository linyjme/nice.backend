package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"niceBackend/common/global"
	"niceBackend/model"
	"niceBackend/utils"
	"os"
	"path"
)

//@author: yjLin
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Gorm() *gorm.DB {
	switch global.NICE_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

// MysqlTables
//@author: yjLin
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func SqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		//model.ServerConfig{},
	)
	if err != nil {
		global.NICE_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.NICE_LOG.Info("register table success")
}

//@author: yjLin
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := global.NICE_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         128,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		global.NICE_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		//sqlDB, _ := db.DB()
		//sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		//sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@author: yjLin
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormSqlite() *gorm.DB {
	projectDir := utils.GetProjectDirectory()
	sqlitePath := path.Join(projectDir, "config", "raysync.db")
	if db, err := gorm.Open(sqlite.Open(sqlitePath), gormConfig()); err != nil {
		return nil
	} else {
		return db
	}
}

//@author: yjLin
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig() *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{TablePrefix: "t_", SingularTable: true},
	}
	return config
}
