package env

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

var Config struct {
	DBConnectionString string
	AppHost            string
}

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	Config.DBConnectionString = os.Getenv("DBConnectionString")
	Config.AppHost = os.Getenv("AppHost")
	err := validateConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func validateConfig() error {
	if len(Config.DBConnectionString) == 0 {
		return NewAppConfigError("Invalid env for database connection string")
	}

	if len(Config.AppHost) == 0 {
		return NewAppConfigError("Invalid host")
	}
	return nil
}
