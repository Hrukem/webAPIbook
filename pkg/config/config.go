package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type C struct {
	Port      string
	Dbuser    string
	Dbpass    string
	Dbname    string
	Dbhost    string
	Dbport    string
	SecretJWT string
}

var Cfg C

// Config function implements configuration
func Config() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("error config.Config()")
		return err
	}
	Cfg.SecretJWT = os.Getenv("SECRET_JWT")

	Cfg.Port = ":4004"
	Cfg.Dbuser = "postgres"
	Cfg.Dbpass = "postgres"
	Cfg.Dbname = "webapibooks"
	Cfg.Dbhost = "dbwebapibooks"
	Cfg.Dbport = "5432"

	return nil
}
