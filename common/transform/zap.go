package transform

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
