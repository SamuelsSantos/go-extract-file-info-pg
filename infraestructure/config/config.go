package config

// Config struct
type Config struct {
	Db     *DbConfig
	Server *ServerConfig
}

//NewConfig new struct to configurations enviroments
func NewConfig() *Config {

	return &Config{
		Db:     NewDbConfig(),
		Server: NewServerConfig(),
	}
}
