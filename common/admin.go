package common

type Admin struct {
	admin_host          string `mapstructure:"admin_host" json:"level" ini:"admin_host"`            //
	admin_port          int `mapstructure:"admin_port" json:"format" ini:"admin_port"`         //
	admin_https_port    int `mapstructure:"admin_https_port" json:"prefix" ini:"admin_https_port"`         //
}