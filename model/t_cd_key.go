package model

import (
	"GameMall/dal"
	"time"
)

type TCdKey struct {
	ID         int64     `gorm:"column:id" json:"id"`
	CdKey      string    `gorm:"column:cd_key" json:"cd_key"`
	Amount     int32     `gorm:"column:amount" json:"amount"`
	IsDelete   int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TCdKey) TableName() string {
	return "t_cd_key"
}

func InsertTCdKey(cdKey string, amount int32) error {
	err := dal.EduDB.Create(&TCdKey{
		CdKey:  cdKey,
		Amount: amount,
	}).Error
	return err
}

func GetTCdKey(cdKey string) (TCdKey, error) {
	var res TCdKey
	err := dal.EduDB.Where("cd_key = ? AND is_delete = 0", cdKey).Find(&res).Error
	return res, err
}

func DelTCdKey(cdKey string) error {
	err := dal.EduDB.Table("t_cd_key").Where("cd_key = ?", cdKey).Update("is_delete", 1).Error
	return err
}
