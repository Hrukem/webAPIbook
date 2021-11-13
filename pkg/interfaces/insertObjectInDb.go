package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func () InsertObjectInDb(r *http.Request) (map[string]int, error){
	var b map[string]string
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Println("error decode body in transport.bookCreate()", err)
		return nil, err
	}

	id, err := db.insert(b)

	m := map[string]int{"id": id}

	return m, nil
}