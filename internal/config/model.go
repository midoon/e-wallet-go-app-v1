package config

type Config struct {
	Server   Server
	Database Database
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
