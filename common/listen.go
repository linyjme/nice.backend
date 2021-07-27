package common

type Listen struct {
	host          string `mapstructure:"host" json:"level" ini:"host"`                      //
	port          int    `mapstructure:"port" json:"format" ini:"port"`                     //
	https_port    int    `mapstructure:"https_port" json:"prefix" ini:"https_port"`         //
	server5port   int    `mapstructure:"server5port" json:"director"  ini:"server5port"`    //
	munu_port     int    `mapstructure:"munu_port" json:"linkName" ini:"munu_port"`         //
	munu_web_port int    `mapstructure:"munu_web_port" json:"showLine" ini:"munu_web_port"` //
}
