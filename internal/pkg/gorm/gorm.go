package gorm

import (
	"niceBackend/config"
	"niceBackend/pkg/util"
	"os"
	"path"

	"niceBackend/common/global"
	"niceBackend/internal/repository/db_repo/admin_repo"
	"niceBackend/internal/repository/db_repo/announcement_repo"
	"niceBackend/internal/repository/db_repo/article_repo"
	"niceBackend/internal/repository/db_repo/categroy_repo"
	"niceBackend/internal/repository/db_repo/comment_repo"
	"niceBackend/internal/repository/db_repo/tag_repo"
	"niceBackend/source"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//@author: yjLin
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Gorm() *gorm.DB {
	switch config.GetConf().System.DbType {
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
		tag_repo.Tags{},
		categroy_repo.Category{},
		announcement_repo.Announcement{},
		comment_repo.Comments{},
		admin_repo.Admin{},
		article_repo.Article{},
	)
	if err != nil {
		global.NiceLog.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	if config.GetConf().System.DbMigrate == true {
		err = source.InitDB(source.Admin)
		if err != nil {
			global.NiceLog.Error("init table failed", zap.Any("err", err))
		} else {
			config.GetConf().System.DbMigrate = false
			config.SetSystemConfig(config.GetConf().System)
		}
	}

	global.NiceLog.Info("register table success")
}

//@author: yjLin
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := config.GetConf().Mysql.Write
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
		global.NiceLog.Error("MySQL启动异常", zap.Any("err", err))
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
	projectDir := util.GetProjectDirectory()
	sqlitePath := path.Join(projectDir, "config", "nice.db")
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
		NamingStrategy:                           schema.NamingStrategy{TablePrefix: "", SingularTable: true},
	}
	return config
}
