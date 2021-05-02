package service

import (
	"GameMall/logs"
	"GameMall/model"
)

func GetKeyAmount(cdKey string) (amount int32) {
	key, err := model.GetTCdKey(cdKey)
	if err != nil {
		logs.Logger.Errorf("func:%v, err:%v", "model.GetTCdKey", err)
		return
	}
	return key.Amount
}
