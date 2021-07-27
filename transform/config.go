package transform

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" ini:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" ini:"zap"`
	Timer   Timer   `mapstructure:"timer" json:"timer" ini:"timer"`
	Listen  Listen  `mapstructure:"listen" json:"listen" ini:"listen"`
	Admin   Admin   `mapstructure:"admin" json:"admin" ini:"admin"`
	RaySync RaySync `mapstructure:"config" json:"config" ini:"config"`
}
