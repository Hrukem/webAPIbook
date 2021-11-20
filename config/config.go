package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gosidekick/goconfig"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type C struct {
	Port   string
	Dbuser string
	Dbpass string
	Dbname string
	Dbhost string
	Dbport string
}

// Cfg new struct config
var Cfg C

// Config function implements configuration
func Config() error {
	vv := C{}
	goconf := C{}
	err := envconfig.Process("w", &vv)
	err = envconfig.Process("w", &Cfg)
	if err != nil {
		log.Println("error read Config")
		return err
	}
	err = goconfig.Parse(&goconf)
	fmt.Println(err)
	fmt.Println("Port", goconf.Port)
	fmt.Println("DbName ", goconf.Dbname)
	_, err = fmt.Printf("Port %s \n", Cfg.Port)
	fmt.Println("err ", err)

	fmt.Println("config :", vv.Port, vv.Dbname, Cfg.Port, Cfg.Dbname)
	fmt.Println("end config")

	envcon := C{}
	err = env.Parse(&envcon)
	fmt.Println(err)
	fmt.Println("envcon :", envcon)

	Cfg.Port = ":4004"
	Cfg.Dbuser = "postgres"
	Cfg.Dbpass = "postgres"
	Cfg.Dbname = "webapibooks"
	Cfg.Dbhost = "dbwebapibooks"
	Cfg.Dbport = "5432"

	return nil
}
