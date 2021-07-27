package transform

type Admin struct {
	AdminHost      string `mapstructure:"admin_host" json:"level" ini:"admin_host"`              //
	AdminPort      int    `mapstructure:"admin_port" json:"format" ini:"admin_port"`             //
	AdminHttpsPort int    `mapstructure:"admin_https_port" json:"prefix" ini:"admin_https_port"` //
}
