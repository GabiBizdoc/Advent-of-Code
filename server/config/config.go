package env

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
	"strings"
)

var Config struct {
	DBConnectionString string
	AppHost            string
	IsDev              bool
}

func LoadConfig() {
	Config.DBConnectionString = os.Getenv("DBConnectionString")
	Config.AppHost = os.Getenv("AppHost")
	err := validateConfig()
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(Config.AppHost, "localhost") {
		Config.IsDev = true
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
