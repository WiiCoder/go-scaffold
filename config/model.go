package config

// BaseConfig 基础配置
type BaseConfig struct {
	Name string `mapstructure:"name"` // 服务名
	Port int32  `mapstructure:"port"` // 服务端口
}

// AuthConfig 通用认证信息
type AuthConfig struct {
	Username string `mapstructure:"username"` // 用户名
	Password string `mapstructure:"password"` // 密码
}

type MysqlConfig struct {
	Host string     `mapstructure:"host"` // 地址
	Port int32      `mapstructure:"port"` // 端口
	Ds   string     `mapstructure:"ds"`   // database
	Auth AuthConfig `mapstructure:"auth"` // 认证信息
}

// EsConfig Es配置
type EsConfig struct {
	Host string     `mapstructure:"host"` // 服务地址
	Post int32      `mapstructure:"post"` // 服务端口
	Auth AuthConfig `mapstructure:"auth"` // 认证信息
}

// TelegramConfig tg配置
type TelegramConfig struct {
	Api   string `mapstructure:"api"`
	Token string `mapstructure:"token"`
	Proxy string `mapstructure:"proxy"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Post     int    `mapstructure:"post"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// ServerConfig 根配置
type ServerConfig struct {
	Base        BaseConfig     `mapstructure:"server"`
	EsConfig    EsConfig       `mapstructure:"elastic"`
	MysqlConfig MysqlConfig    `mapstructure:"mysql"`
	TgConfig    TelegramConfig `mapstructure:"telegram"`
	RedisConfig RedisConfig    `mapstructure:"redis"`
}
