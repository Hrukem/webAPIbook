package process

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage"
	"log"
	"net/http"
)

func (bs *Proc) InsertObjectInDb(r *http.Request, db *storage.DB) ([]byte, error) {
	var b storage.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Println("error decode body in process.InsertObjectInDb()", err)
		return nil, err
	}

	id, err := bs.Insert(b, db)

	m := map[string]int{"id": id}

	answer, err := json.Marshal(m)
	if err != nil {
		log.Println("error Marshal in process.InsertObjectInDb()", err)
		return nil, err
	}

	return answer, nil
}
