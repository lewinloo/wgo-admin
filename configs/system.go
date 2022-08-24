package config

type System struct {
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                            // 端口值
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                   // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	GlobalPrefix string `mapstructure:"global-prefix" json:"global-prefix" yaml:"global-prefix"` // 全局前缀
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`             // 是否使用redis
	UseCron      bool   `mapstructure:"use-cron" json:"use-cron" yaml:"use-cron"`                // 是否使用cron
}
