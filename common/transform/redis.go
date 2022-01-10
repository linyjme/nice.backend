package transform

type Redis struct {
	DB           int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Pass         string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Db           int    `toml:"db"`
	MaxRetries   int    `toml:"maxRetries"`
	PoolSize     int    `toml:"poolSize"`
	MinIdleConns int    `toml:"minIdleConns"`
}
