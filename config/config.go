package config

import (
	"fmt"
	"os"
)

const (
	// EnvDev const represents dev environment
	EnvDev = "dev"
	// EnvStaging const represents staging environment
	EnvStaging = "staging"
	// EnvProduction const represents production environment
	EnvProduction = "production"
)

// Env holds the current environment
var (
	Env          string
	Port         string
	DBDriver     string
	DBDataSource string
)

// TEST Credentials ------------------------------------------------------------
var (
	TestDBDriver     = "postgres"
	TestDBDataSource = "user=postgres password=postgres dbname=postgres sslmode=disable host=localhost"
	// docker run -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=postgres postgres
)

// Initialize ...
func Initialize() {
	GetAllEnv()
}

// GetAllEnv should get all the env configs required for the app.
func GetAllEnv() {
	// API Configs
	mustEnv("ENV", &Env, EnvDev)
	mustEnv("PORT", &Port, "8080")
	mustEnv("DB_DRIVER", &DBDriver, "postgres")
	mustEnv("DB_DATASOURCE", &DBDataSource,
		"user=postgres password=postgres dbname=weather_monster sslmode=disable host=localhost")
}

// mustEnv get the env variable with the name 'key' and store it in 'value'
func mustEnv(key string, value *string, defaultVal string) {
	if *value = os.Getenv(key); *value == "" {
		*value = defaultVal
		fmt.Printf("%s env variable not set, using default value.\n", key)
	}
}
