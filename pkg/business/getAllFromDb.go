package business

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage"
)

func (bs *B) GetAllFromDb(db *storage.DB) ([]byte, error) {
	sliceObjects, err := bs.SelectAll(db)
	if err != nil {
		return nil, err
	}

	sliceByte, err := json.Marshal(sliceObjects)
	if err != nil {
		fmt.Println("error Marshal in interfaces.getAll()")
		return nil, err
	}

	return sliceByte, nil
}