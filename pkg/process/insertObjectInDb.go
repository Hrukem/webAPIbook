package process

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage/postgress"
	"golang_ninja/webAPIbook/pkg/transport/g_r_p_c/grpcServer"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func (p *Proc) InsertObjectInDb(
	r *http.Request,
	db *postgress.DB,
	loggingInMongo chan<- grpcServer.L,
) ([]byte, error) {
	var b postgress.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Println("error decode body in process.InsertObjectInDb()", err)
		return nil, err
	}

	id, err := p.Insert(b, db)

	m := map[string]int{"id": id}

	answer, err := json.Marshal(m)
	if err != nil {
		log.Println("error Marshal in process.InsertObjectInDb()", err)
		return nil, err
	}

	// logging in MongoDb
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 300
	idUser := rand.Intn(max-min+1) + min
	l := grpcServer.L{Id: int64(idUser), Log: "getObjectsFromDb"}
	loggingInMongo <- l

	return answer, nil
}
