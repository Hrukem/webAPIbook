package process

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage"
)

func (p *Proc) GetAllFromDb(db *storage.DB) ([]byte, error) {
	sliceObjects, err := p.SelectAll(db)
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
