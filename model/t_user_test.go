package model

import (
	"GameMall/dal"
	"GameMall/tools"
	"fmt"
	"testing"
)

func init() {
	err := dal.InitDB()
	if err != nil {
		panic(err)
	}
}

func TestGetTUserInfo(t *testing.T) {
	info, err := GetTUserInfo(&TUser{Id: 10001})
	if err != nil {
		panic(err)
	}
	fmt.Println(len(info))
}

func TestInsertTUser(t *testing.T) {
	u, err := InsertTUser(&TUser{Account: "15200830961", Nickname: "MemoKun2", Password: "123"})
	if err != nil {
		panic(err)
	}
	fmt.Println(tools.ToJson(u))
}

func TestIncreaseBalance(t *testing.T) {
	err := IncreaseBalance(10003, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok!")
}
