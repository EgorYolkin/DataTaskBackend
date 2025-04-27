package config

// Config interface
type Config struct {
	HTTP     HTTP
	Database Database
	RabbitMQ RabbitMQ
	Swagger  Swagger
}

type HTTP struct {
	Host          string   `mapstructure:"host"`
	Port          int      `mapstructure:"port"`
	Mode          string   `mapstructure:"mode"`
	AllowOrigins  []string `mapstructure:"allow_origins"`
	AllowHeaders  []string `mapstructure:"allow_headers"`
	AllowMethods  []string `mapstructure:"allow_methods"`
	ExposeHeaders []string `mapstructure:"expose_headers"`
}

type Swagger struct {
	BasePath string `mapstructure:"base_path"`
	Version  string `mapstructure:"version"`
}

type Database struct {
	Host string `mapstructure:"db_host" json:"db_host"`
	Port string `mapstructure:"db_port" json:"db_port"`
	User string `mapstructure:"db_user" json:"db_user"`
	Pass string `mapstructure:"db_pass" json:"db_pass"`
	Base string `mapstructure:"db_base" json:"db_base"`
}

type RabbitMQ struct {
	Host string `mapstructure:"rabbitmq_host" json:"rabbitmq_host"`
	Port string `mapstructure:"rabbitmq_port" json:"rabbitmq_port"`
	User string `mapstructure:"rabbitmq_user" json:"rabbitmq_user"`
	Pass string `mapstructure:"rabbitmq_pass" json:"rabbitmq_pass"`
}
