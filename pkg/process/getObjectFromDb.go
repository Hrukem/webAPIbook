package process

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage/postgress"
	"golang_ninja/webAPIbook/pkg/transport/g_r_p_c/grpcServer"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func (p *Proc) GetObjectFromDb(
	db *postgress.DB,
	path string,
	loggingInMongo chan<- grpcServer.L,
) ([]byte, error) {
	match, err := regexp.MatchString(`/book/[\d]+$`, path)
	if !match {
		log.Println("error match path in process.GetObjectFromDb()")
		return nil, err
	}

	re, err := regexp.Compile(`[\d]+`)
	if err != nil {
		log.Println("error regexp in process.GetObjectFromDb()")
		return nil, err
	}
	idString := re.FindString(path)

	// from cash
	object := db.Get(idString)
	if object != nil {
		sliceByte, err := json.Marshal(object)
		if err != nil {
			fmt.Println("error Marshal in interfaces.getAll()")
			return nil, err
		}

		// logging in MongoDb
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 300
		idUser := rand.Intn(max-min+1) + min
		l := grpcServer.L{Id: int64(idUser), Log: "getObjectsFromDb"}
		loggingInMongo <- l

		return sliceByte, nil
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("error conv string in process.GetObjectFromDb()")
		return nil, err
	}

	// from postgresDB
	object, err = p.SelectObject(db, id)
	if err != nil {
		log.Println("error select object from Db in process.GetObjectFromDb()")
		return nil, err
	}

	sliceByte, err := json.Marshal(object)
	if err != nil {
		fmt.Println("error Marshal in interfaces.getAll()")
		return nil, err
	}

	// logging in MongoDb
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 300
	idUser := rand.Intn(max-min+1) + min
	l := grpcServer.L{Id: int64(idUser), Log: "getObjectsFromDb"}
	loggingInMongo <- l

	return sliceByte, nil
}
