package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppDebug   bool
<<<<<<< HEAD
	DBMigrate  bool
	DBSeeding  bool
=======
>>>>>>> da0e613 (Initial commit)
}

func LoadConfig() (Config, error) {

	localEnv := viper.New()
	localEnv.SetConfigType("dotenv")
	viper.SetConfigFile("../.env") // Specify the config file name

	// Set default values
	viper.SetDefault("DBHost", "localhost")
	viper.SetDefault("DBPort", "5432")
	viper.SetDefault("DBUser", "user")
	viper.SetDefault("DBPassword", "password")
	viper.SetDefault("DBName", "database")
	viper.SetDefault("AppDebug", true)

<<<<<<< HEAD
	viper.SetDefault("DBMigrate", false)
	viper.SetDefault("DBSeeding", false)

=======
>>>>>>> da0e613 (Initial commit)
	// Allow Viper to read environment variables
	viper.AutomaticEnv()

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %s, using default values or environment variables", err)
	}

	// add value to the config
	config := Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		AppDebug:   viper.GetBool("APP_DEBUG"),
<<<<<<< HEAD
		DBMigrate:  viper.GetBool("DB_MIGRATE"),
		DBSeeding:  viper.GetBool("DB_SEEDING"),
=======
>>>>>>> da0e613 (Initial commit)
	}

	return config, nil
}
