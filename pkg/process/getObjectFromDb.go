package process

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage"
	"log"
	"regexp"
	"strconv"
)

func (p *Proc) GetObjectFromDb(db *storage.DB, path string) ([]byte, error) {
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

	object := db.Get(idString)
	if object != nil {
		sliceByte, err := json.Marshal(object)
		if err != nil {
			fmt.Println("error Marshal in interfaces.getAll()")
			return nil, err
		}
		return sliceByte, nil
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("error conv string in process.GetObjectFromDb()")
		return nil, err
	}

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

	return sliceByte, nil
}
