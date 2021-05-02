package model

import (
	"GameMall/common"
	"GameMall/dal"
	"GameMall/tools"
	"errors"
	"time"
)

type TProduct struct {
	Id            int64     `gorm:"column:id" json:"id"`
	ProviderId    int64     `gorm:"column:provider_id" json:"provider_id"`
	ProviderName  string    `gorm:"column:provider_name" json:"provider_name"`
	Name          string    `gorm:"column:name" json:"name"`
	ProductType   int8      `gorm:"column:product_type" json:"product_type"`
	Status        int8      `gorm:"column:status" json:"status"`
	Price         int32     `gorm:"column:price" json:"price"`
	Keywords      string    `gorm:"column:keywords" json:"keywords"`
	AfterSaleText string    `gorm:"column:after_sale_text" json:"after_sale_text"`
	Inventory     int32     `gorm:"column:inventory" json:"inventory"`
	SaleVolume    int32     `gorm:"column:sale_volume" json:"sale_volume"`
	DescText      string    `gorm:"column:desc_text" json:"desc_text"`
	DescImg       string    `gorm:"column:desc_img" json:"desc_img"`
	IsDelete      int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TProduct) TableName() string {
	return "t_product"
}

func GetTProduct(cond *TProduct) ([]*TProduct, error) {
	var res []*TProduct
	err := dal.EduDB.Where(cond).Where(map[string]interface{}{"is_delete": 0}).Find(&res).Error
	return res, err
}

func QueryTProduct(cond *TProduct, pageNum, pageSize int32) ([]*TProduct, int64, error) {
	if cond == nil {
		return nil, 0, errors.New("query no cond")
	}

	var res []*TProduct
	tx := dal.EduDB.Table("t_product").Where("is_delete = 0")
	if cond.ProviderName != "" {
		tx = tx.Where("provider_name like ?", "%"+cond.ProviderName+"%")
	}
	if cond.Keywords != "" {
		tx = tx.Where("keywords like ? OR name like ?", "%"+cond.Keywords+"%", "%"+cond.Keywords+"%")
	}
	if tools.InArrayNoIndex(cond.ProductType, []int8{common.ProductTypeExperience, common.ProductTypeRegular}) {
		tx = tx.Where("product_type = ?", cond.ProductType)
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

func InsertTProduct(data *TProduct) (*TProduct, error) {
	if data == nil {
		return nil, errors.New("insert no data")
	}
	err := dal.EduDB.Create(data).Error
	return data, err
}

func UpdateTProduct(data *TProduct) error {
	if data == nil {
		return errors.New("update no data")
	}
	err := dal.EduDB.Where("id = ?", data.Id).Update(data).Error
	return err
}
