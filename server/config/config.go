package config

type Config struct {
	Mysql       *Mysql       `mapstructure:"mysql"`
	Redis       *Redis       `mapstructure:"redis"`
	Rabbitmq    *Rabbitmq    `mapstructure:"rabbitmq"`
	DomainName  string       `mapstructure:"domain_name"`
	ServerPort  int          `mapstructure:"server_port"`
	Debug       bool         `mapstructure:"debug"`
	Key         string       `mapstructure:"key"`
	AdminConfig *AdminConfig `mapstructure:"admin_config"`
}
