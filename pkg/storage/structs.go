package storage

type book struct {
	Title      string
	Author     string
	Publishing string
}

type books struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publishing string `json:"publishing"`
}
