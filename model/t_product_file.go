package model

import (
	"GameMall/dal"
	"errors"
	"time"
)

type TProductFile struct {
	Id         int64     `gorm:"column:id" json:"id"`
	ProductID  int64     `gorm:"column:product_id" json:"product_id"`
	FileType   int8      `gorm:"column:file_type" json:"file_type"`
	FileName   string    `gorm:"column:file_name" json:"file_name"`
	Url        string    `gorm:"column:url" json:"url"`
	IsDelete   int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TProductFile) TableName() string {
	return "t_product_file"
}

func GetTProductFile(cond *TProductFile) ([]*TProductFile, error) {
	var res []*TProductFile
	err := dal.EduDB.Where(cond).Where(map[string]interface{}{"is_delete": 0}).Find(&res).Error
	return res, err
}

func InsertTProductFile(data *TProductFile) (*TProductFile, error) {
	if data == nil {
		return nil, errors.New("insert no data")
	}
	err := dal.EduDB.Create(data).Error
	return data, err
}

func UpdateTProductFile(data *TProductFile) error {
	if data == nil {
		return errors.New("update no data")
	}
	err := dal.EduDB.Where("id = ?", data.Id).Update(data).Error
	return err
}
