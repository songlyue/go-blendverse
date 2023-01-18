package config

type ServerConfig struct {
	Name       string `mapstructure:"name" json:"name" yaml:"name"`
	Port       string `mapstructure:"port" json:"port" yaml:"port"`
	LogAddress string `mapstructure:"logAddress" json:"logAddress" yaml:"logAddress"`
}

type MysqlConfig struct {
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DbName   string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
	Port     int
}

type PgsqlConfig struct {
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DbName   string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}

type LogConfig struct {
	LogFilePath     string `mapstructure:"logFilePath" json:"logFilePath" yaml:"logFilePath"`
	LogInfoFileName string `mapstructure:"logInfoFileName" json:"logInfoFileName" yaml:"logInfoFileName"`
	LogWarnFileName string `mapstructure:"logWarnFileName" json:"logWarnFileName" yaml:"logWarnFileName"`
	LogFileExt      string `mapstructure:"logFileExt" json:"logFileExt" yaml:"logFileExt"`
}
