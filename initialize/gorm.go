package initialize

import (
	"asyncClient/common/global"
	"asyncClient/model"
	"asyncClient/utils"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path"
)

//@author: linyj
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Gorm() *gorm.DB {
	switch global.RAY_CONFIG.RaySync.UseMysql {
	case "1":
		return GormMysql()
	case "0":
		return GormSqlite()
	default:
		return GormSqlite()
	}
}

// MysqlTables
//@author: linyj
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func SqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		global.RAY_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.RAY_LOG.Info("register table success")
}

//@author: linyj
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := global.RAY_CONFIG.RaySync
	if m.UseMysql == "0" {
		return nil
	}
	dsn := m.MysqlUser + ":" + m.MysqlPassword + "@tcp(" + m.MysqlHost + ":" + m.MysqlPort + ")/" + m.MysqlName + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		global.RAY_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		//sqlDB, _ := db.DB()
		//sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		//sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@author: linyj
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

//@author: linyj
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig() *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{TablePrefix: "tb_", SingularTable: true},
	}
	return config
}
