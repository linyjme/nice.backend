package transform

type Server struct {
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Timer    Timer    `mapstructure:"timer" json:"timer" yaml:"timer"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mongo    Mongodb  `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb"`
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Language Language `mapstructure:"language" json:"language" yaml:"language"`
	Mail     Mail     `mapstructure:"mail" json:"mail" ini:"mail"`
}
