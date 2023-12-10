package env

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
)

var Config struct {
	DBConnectionString string
	AppHost            string
}

func LoadConfig() {
	Config.DBConnectionString = os.Getenv("DBConnectionString")
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
