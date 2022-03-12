package config

import "time"

type Server struct {
	JWT    JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Timer  Timer   `mapstructure:"timer" json:"timer" yaml:"timer"`
	System System  `mapstructure:"system" json:"system" yaml:"system"`
	Redis  Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mongo  Mongodb `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb"`
	Mysql  Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mail   Mail    `mapstructure:"email" json:"email" yaml:"email"`
}

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
}

type Mysql struct {
	Read struct {
		Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 服务器地址:端口
		Config   string `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
		Dbname   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`     // 数据库名
		Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
		Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	} `yaml:"read"`
	Write struct {
		Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 服务器地址:端口
		Config   string `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
		Dbname   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`     // 数据库名
		Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
		Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	} `yaml:"write"`
	Base struct {
		MaxIdleConns    int           `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
		MaxOpenConns    int           `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
		LogMode         string        `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // 是否开启Gorm全局日志
		LogZap          bool          `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     // 是否通过zap写入日志文件
		ConnMaxLifeTime time.Duration `mapstructure:"connMaxLifeTime" json:"connMaxLifeTime" yaml:"connMaxLifeTime"`
	} `yaml:"base"`
}

//func (m *Mysql) Dsn() string {
//	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
//}

type Redis struct {
	DB           int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Pass         string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	MaxRetries   int    `toml:"maxRetries"`
	PoolSize     int    `toml:"poolSize"`
	MinIdleConns int    `toml:"minIdleConns"`
}

type Mongodb struct {
	DB       string `mapstructure:"db" json:"db" yaml:"db"`                   // mongo的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Number   uint64 `mapstructure:"number" json:"number" yaml:"number"`       // 连接是
}

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // 环境值
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`                              // 端口值
	Domain        string `mapstructure:"domain" json:"domain" yaml:"domain"`                        // 域名
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
	DbMigrate     bool   `mapstructure:"db_migrate" json:"db_migrate" yaml:"db_migrate"`            // 初始化数据标志
	Language      string `mapstructure:"language" json:"language" yaml:"language"`                  // 初始化数据标志
}

type Timer struct {
	Start  bool     `mapstructure:"start" json:"start" yaml:"start"` // 是否启用
	Spec   string   `mapstructure:"spec" json:"spec" yaml:"spec"`    // CRON表达式
	Detail []Detail `mapstructure:"detail" json:"detail" yaml:"detail"`
}

type Detail struct {
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // 需要清理的表名
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // 需要比较时间的字段
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // 时间间隔
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" ini:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" ini:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" ini:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  ini:"director"`                 // 日志文件夹
	LinkName      string `mapstructure:"link-name" json:"linkName" ini:"link-name"`                // 软链接名称
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" ini:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" ini:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" ini:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" ini:"log-in-console"`  // 输出控制台
}

type Mail struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
	To   string `toml:"to"`
}
