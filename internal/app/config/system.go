package config

type System struct {
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                            // 端口值
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                   // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	GlobalPrefix string `mapstructure:"global-prefix" json:"global-prefix" yaml:"global-prefix"` // 全局前缀
	Mode         string `mapstructure:"mode" json:"mode" yaml:"mode"`                            // 系统模式
}
