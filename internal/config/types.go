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
	Host string `mapstructure:"DB_HOST"`
	Port string `mapstructure:"DB_PORT"`
	User string `mapstructure:"DB_USER"`
	Pass string `mapstructure:"DB_PASS"`
	Base string `mapstructure:"DB_BASE"`
}

type RabbitMQ struct {
	Host string `mapstructure:"RABBITMQ_HOST"`
	Port string `mapstructure:"RABBITMQ_PORT"`
	User string `mapstructure:"RABBITMQ_USER"`
	Pass string `mapstructure:"RABBITMQ_PASS"`
}

type JWT struct {
	Secret string `mapstructure:"JWT_SECRET"`
}
