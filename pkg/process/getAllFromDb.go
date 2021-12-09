package process

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage/postgress"
	"golang_ninja/webAPIbook/pkg/transport/g_r_p_c/grpcServer"
	"math/rand"
	"time"
)

func (p *Proc) GetAllFromDb(
	db *postgress.DB,
	loggingInMongo chan<- grpcServer.L,
) ([]byte, error) {
	sliceObjects, err := p.SelectAll(db)
	if err != nil {
		return nil, err
	}

	sliceByte, err := json.Marshal(sliceObjects)
	if err != nil {
		fmt.Println("error Marshal in interfaces.getAll()")
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 300
	idUser := rand.Intn(max-min+1) + min

	l := grpcServer.L{Id: int64(idUser), Log: "getAllObjects"}

	loggingInMongo <- l
	return sliceByte, nil
}
