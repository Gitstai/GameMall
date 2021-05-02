package model

import (
	"GameMall/dal"
	"errors"
	"time"
)

type TPurchaseRecords struct {
	Id            int64     `gorm:"column:id" json:"id"`
	UserId        int64     `gorm:"column:user_id" json:"user_id"`
	ProductId     int64     `gorm:"column:product_id" json:"product_id"`
	Payment       int32     `gorm:"column:payment" json:"payment"`
	ProductDesc   string    `gorm:"column:product_desc" json:"product_desc"`
	ProductName   string    `gorm:"column:product_name" json:"product_name"`
	ProductType   int8      `gorm:"column:product_type" json:"product_type"`
	DescImg       string    `gorm:"column:desc_img" json:"desc_img"`
	AfterSaleText string    `gorm:"column:after_sale_text" json:"after_sale_text"`
	IsDelete      int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TPurchaseRecords) TableName() string {
	return "t_purchase_records"
}

func GetTPurchaseRecords(cond *TPurchaseRecords) ([]*TPurchaseRecords, error) {
	var res []*TPurchaseRecords
	err := dal.EduDB.Where(cond).Where(map[string]interface{}{"is_delete": 0}).Find(&res).Error
	return res, err
}

func InsertTPurchaseRecords(data *TPurchaseRecords) (*TPurchaseRecords, error) {
	if data == nil {
		return nil, errors.New("insert no data")
	}
	err := dal.EduDB.Create(data).Error
	return data, err
}

func QueryTPurchaseRecords(cond *TPurchaseRecords, pageNum, pageSize int32) ([]*TPurchaseRecords, int64, error) {
	if cond == nil {
		return nil, 0, errors.New("query no cond")
	}

	var res []*TPurchaseRecords
	tx := dal.EduDB.Table("t_purchase_records").Where("user_id = ? AND is_delete = 0", cond.UserId)

	if cond.ProductName != "" {
		tx = tx.Where("product_name like ?", "%"+cond.ProductName+"%")
	}

	var count int64
	err := tx.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	if pageNum > 0 && pageSize > 0 {
		tx = tx.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}

	err = tx.Find(&res).Error
	return res, count, err
}
