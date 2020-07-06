package config

import "os"

// Config struct
type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	DbPassword string
	DbDriver   string
	ServerPort string
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

//GetConf read configurations
func GetConf() Config {

	return Config{
		DbHost:     getenv("DB_HOST", "localhost"),
		DbPort:     getenv("DB_PORT", "5432"),
		DbUser:     getenv("DB_USER", "postgres"),
		DbName:     getenv("DB_NAME", "import-data"),
		DbPassword: getenv("DB_PASSWORD", "db@123A"),
		DbDriver:   getenv("DB_DRIVER", "postgres"),
		ServerPort: getenv("SERVER_PORT", "8085"),
	}
}
