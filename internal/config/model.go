package config

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type JWT struct {
	Key    string
	Issuer string
}

type Redis struct {
	Addr     string
	Password string
}

type RabbitMQ struct {
	Username string
	Password string
	Host     string
	Port     string
	User     string
	Exchange string
	RKey     string
	Queue    string
}
