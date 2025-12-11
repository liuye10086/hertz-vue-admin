package config

type Server struct {
	OnlineUser OnlineUser `mapstructure:"online-user" json:"online-user" yaml:"online-user"`
	JWT        JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap        Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email      Email      `mapstructure:"email" json:"email" yaml:"email"`
	System     System     `mapstructure:"system" json:"system" yaml:"system"`
	Captcha    Captcha    `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
