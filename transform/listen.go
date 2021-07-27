package transform

type Listen struct {
	Host        string `mapstructure:"host" json:"level" ini:"host"`                      //
	Port        int    `mapstructure:"port" json:"format" ini:"port"`                     //
	HttpsPort   int    `mapstructure:"https_port" json:"prefix" ini:"https_port"`         //
	Server5port int    `mapstructure:"server5port" json:"director"  ini:"server5port"`    //
	MunuPort    int    `mapstructure:"munu_port" json:"linkName" ini:"munu_port"`         //
	MunuWebPort int    `mapstructure:"munu_web_port" json:"showLine" ini:"munu_web_port"` //
}
