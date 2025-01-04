package bootstrap

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	DBUser   string `mapstructure:"DB_USER"`
	DBPass   string `mapstructure:"DB_PASS"`
	DBHost   string `mapstructure:"DB_HOST"`
	DBPort   int    `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_NAME"`
	SSLMode  string `mapstructure:"SSL_MODE"`
	TimeZone string `mapstructure:"TIME_ZONE"`
}

func NewEnv() *Env {
	env := Env{}
	if len(os.Args) == 2 {
		if os.Args[1] == "dev" {
			viper.AddConfigPath(".")
			viper.SetConfigName(".env.local")
			viper.SetConfigType("env")
		} else if os.Args[1] == "prod" {
			viper.AddConfigPath(".")
			viper.SetConfigName(".env.prod")
			viper.SetConfigType("env")
		}
	} else {
		viper.SetConfigFile(".env")
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &env
}
