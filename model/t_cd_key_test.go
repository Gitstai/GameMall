package model

import (
	"fmt"
	"testing"
)

func TestInsertTCdKey(t *testing.T) {
	err := InsertTCdKey("3901170209", 50000)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
}
