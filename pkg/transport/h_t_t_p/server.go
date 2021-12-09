package h_t_t_p

import (
	httpSwagger "github.com/swaggo/http-swagger"
	_ "golang_ninja/webAPIbook/cmd/webAPIbook/docs" // docs are generated by Swag CLI, you have to import it.
	"golang_ninja/webAPIbook/pkg/config"
	"golang_ninja/webAPIbook/pkg/process"
	"golang_ninja/webAPIbook/pkg/storage/postgress"
	"golang_ninja/webAPIbook/pkg/transport/authentification"
	"golang_ninja/webAPIbook/pkg/transport/g_r_p_c/grpcServer"
	"log"
	"net/http"
)

type T struct {
	DbPostgres *postgress.DB
	Trnsprt
	process.Proc
	LoggingInMongo chan<- grpcServer.L
}

// Server function start grpcServer
func Server(dbPostgres *postgress.DB, loggingInMongo chan<- grpcServer.L) error {
	t := &T{dbPostgres, Trnsprt{}, process.Proc{}, loggingInMongo}

	rout := http.NewServeMux()

	rout.HandleFunc("/auth/", authentification.GenerationJWT)

	rout.HandleFunc("/books", authentification.CheckAuth(t.GetAll))
	rout.HandleFunc("/book", authentification.CheckAuth(t.Post))
	rout.HandleFunc("/book/", authentification.CheckAuth(t.MethodSwitch))
	rout.HandleFunc(
		"/swagger/",
		httpSwagger.Handler(
			httpSwagger.URL(
				"h_t_t_p://localhost:4004/swagger/doc.json",
			),
		),
	)

	log.Println("start http grpcServer on port", config.Cfg.Port)
	err := http.ListenAndServe(config.Cfg.Port, rout)
	if err != nil {
		log.Println("error grpcServer", err)
	}
	return nil
}
