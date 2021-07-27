package transform

type RaySync struct {
	AuthProtocol     string `mapstructure:"auth_protocol" json:"use_mysql" ini:"auth_protocol"`                    //
	AuthHost         string `mapstructure:"auth_host" json:"auth_host" ini:"auth_host"`                            //
	AuthPort         int    `mapstructure:"auth_port" json:"auth_port" ini:"auth_port"`                            //
	AuthApi          string `mapstructure:"auth_api" json:"auth_api" ini:"auth_api"`                               //
	AuthMethod       string `mapstructure:"auth_method" json:"auth_method" ini:"auth_method"`                      //
	License          string `mapstructure:"license" json:"license" ini:"license"`                                  //
	UseMysql         string `mapstructure:"use_mysql" json:"use_mysql" ini:"use_mysql"`                            //
	MysqlHost        string `mapstructure:"mysql_host" json:"mysql_host" ini:"mysql_host"`                         //
	MysqlPort        string `mapstructure:"mysql_port" json:"mysql_port" ini:"mysql_port"`                         //
	MysqlUser        string `mapstructure:"mysql_user" json:"mysql_user" ini:"mysql_user"`                         //
	MysqlPassword    string `mapstructure:"mysql_password" json:"mysql_password" ini:"mysql_password"`             //
	MysqlName        string `mapstructure:"mysql_name" json:"mysql_name" ini:"mysql_name"`                         //
	SecureSwitch     string `mapstructure:"secure_switch" json:"secure_switch" ini:"secure_switch"`                //
	NameWithoutId    string `mapstructure:"name_without_id" json:"name_without_id" ini:"name_without_id"`          //
	SslPort          int    `mapstructure:"ssl_port" json:"ssl_port" ini:"ssl_port"`                               //
	NoSslPort        int    `mapstructure:"no_ssl_port" json:"no_ssl_port" ini:"no_ssl_port"`                      //
	LicenseHost      string `mapstructure:"licensehost" json:"licensehost" ini:"licensehost"`                      //
	LicensePort      int    `mapstructure:"licenseport" json:"licenseport" ini:"licenseport"`                      //
	LicenseCheck     string `mapstructure:"licensecheck" json:"licensecheck" ini:"licensecheck"`                   //
	WebsocketPort    int    `mapstructure:"websocket_port" json:"websocket_port" ini:"websocket_port"`             //
	WebsocketSslPort int    `mapstructure:"websocket_ssl_port" json:"websocket_ssl_port" ini:"websocket_ssl_port"` //
}
