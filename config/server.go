package config

type Server struct {
	Server ServerConfig `mapstructure:"server" json:"server" yaml:"server"`
	Mysql  MysqlConfig  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  PgsqlConfig  `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Redis  RedisConfig  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Log    LogConfig    `mapstructure:"log" json:"log" yaml:"log"`
	JWT    JWT          `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
