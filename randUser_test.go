package randUser

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func  TestMain(m *testing.M)  {
	flag.Parse()
	os.Exit(m.Run())
}

func TestGetUserName(t *testing.T) {
	for i := 0; i < 1000 ; i++ {

		fmt.Println("-----------------")
		fmt.Println(GetUserName())
	}
}
