package h_t_t_p

import (
	"fmt"
	"os/exec"
	"testing"
)

type FakeTrnsprt struct {
}

func (f FakeTrnsprt) GetAllFromDb() ([]byte, error) {
	b := []byte("goher")
	return b, nil
}

func TestGetAll(t *testing.T) {

	fmt.Println("start test")
	i := exec.Command("docker-compose")
	fmt.Println(i)
	err := i.Run()
	if err != nil {
		fmt.Println("err: ", err)
	}

	f := new(FakeTrnsprt)

	fmt.Println("f", f)
}
