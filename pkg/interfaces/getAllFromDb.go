package interfaces

func (db *DB) getAllFromDb() ([]byte, error) {
	res, err := db.selectAll()
	if err != nil {
		return nil, err
	}

}
