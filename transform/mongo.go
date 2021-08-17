package transform

type Mongodb struct {
	DB         string `mapstructure:"db" json:"db" yaml:"db"`                         // mongo的哪个数据库
	Collection int    `mapstructure:"collection" json:"collection" yaml:"collection"` // mongo的哪个数据库
	Addr       string `mapstructure:"addr" json:"addr" yaml:"addr"`                   // 服务器地址:端口
	Password   string `mapstructure:"password" json:"password" yaml:"password"`       // 密码
	Number     uint64 `mapstructure:"number" json:"number" yaml:"number"`             // 连接是
}
