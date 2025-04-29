package config

// Config interface
type Config struct {
	HTTP     HTTP
	Database Database
	RabbitMQ RabbitMQ
	Swagger  Swagger
	JWT      JWT
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
	Host string `mapstructure:"db_host"`
	Port string `mapstructure:"db_port"`
	User string `mapstructure:"db_user"`
	Pass string `mapstructure:"db_pass"`
	Base string `mapstructure:"db_base"`
}

type RabbitMQ struct {
	Host string `mapstructure:"rabbitmq_host"`
	Port string `mapstructure:"rabbitmq_port"`
	User string `mapstructure:"rabbitmq_user"`
	Pass string `mapstructure:"rabbitmq_pass"`
}

type JWT struct {
	Secret string `mapstructure:"jwt_secret"`
}
