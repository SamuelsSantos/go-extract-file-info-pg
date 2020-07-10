package config

import "fmt"

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
	Driver   string
}

type serverConfig struct {
	Port string
}

// Config struct
type Config struct {
	Server *serverConfig
	Db     *dbConfig
}

//NewConfig new struct to configurations enviroments
func NewConfig() *Config {
	return &Config{
		Server: &serverConfig{
			Port: GetenvString("SERVER_PORT", "8085"),
		},
		Db: &dbConfig{
			Host:     GetenvString("DB_HOST", "127.0.0.1"),
			Port:     GetenvString("DB_PORT", "5432"),
			User:     GetenvString("DB_USER", "postgres"),
			Name:     GetenvString("DB_NAME", "import-data"),
			Password: GetenvString("DB_PASSWORD", "db@123A"),
			Driver:   GetenvString("DB_DRIVER", "postgres"),
		},
	}
}

// ToString return string values from enviroments
func (cfg *dbConfig) ToString() string {
	return fmt.Sprintf("Host: %s\nPort:%s\nUser: %s\nDB:%s", cfg.Host, cfg.Port, cfg.User, cfg.Name)
}
