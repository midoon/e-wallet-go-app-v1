package config

type Config struct {
	Server   Server
	Database Database
	JWT      JWT
	Redis    Redis
}

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
