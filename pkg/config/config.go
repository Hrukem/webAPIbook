package config

type cfg struct {
	Port   string
	Table  string
	DbUser string
	DbPass string
	DbName string
	DbHost string
	DbPort string
}

var Cfg cfg // new struct config

// Config function implements configuration
func Config() error {
	Cfg.Port = "0.0.0.0:4004"
	Cfg.DbUser = "postgres"
	Cfg.DbPass = "postgres"
	Cfg.DbName = "webapibooks"
	Cfg.DbHost = "dbwebapibooks"
	Cfg.DbPort = "5432"

	return nil
}

/* for get from .env file

func Config() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("error config")
		return err
	}

	Cfg.Port = os.Getenv("PORT")
	Cfg.Table = os.Getenv("TABLE")
	Cfg.DbUser = os.Getenv("DATABASE_USER")
	Cfg.DbPass = os.Getenv("DATABASE_PASSWORD")
	Cfg.DbName = os.Getenv("DATABASE_NAME")
	Cfg.DbHost = os.Getenv("DATABASE_HOST")
	Cfg.DbPort = os.Getenv("DATABASE_PORT")

	return nil
}
*/
