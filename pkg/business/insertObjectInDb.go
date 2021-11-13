package business

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage"
	"net/http"
)

func (bs *B) InsertObjectInDb(r *http.Request, db *storage.DB) (map[string]int, error) {
	var b map[string]string
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Println("error decode body in transport.bookCreate()", err)
		return nil, err
	}

	id, err := bs.Insert(b, db)

	m := map[string]int{"id": id}

	return m, nil
}
