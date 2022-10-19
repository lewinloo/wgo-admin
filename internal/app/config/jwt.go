package config

type Jwt struct {
	Secret     string `mapstructure:"secret" json:"secret" yaml:"secret"`
	ExpireTime int    `mapstructure:"expire-time" json:"expire-time" yaml:"expire-time"`
	Issuer     string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	Subject    string `mapstructure:"subject" json:"subject" yaml:"subject"`
}
