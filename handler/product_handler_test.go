package handler

import (
	"GameMall/dto"
	"GameMall/tools"
	"fmt"
	"testing"
)

func TestUpsertEduProduct(t *testing.T) {
	req := dto.UpsertEduProductReq{
		Id:            0,
		UserId:        0,
		ProductName:   "测试商品1",
		ProductType:   2,
		Price:         19999,
		FileType:      2,
		FileUrl:       "www.baidu.com",
		Status:        1,
		ProductDesc:   "这个教育视频，很有教育意义！",
		BannerImgs:    []string{"123", "456", "789"},
		AfterSaleText: "购买了本视频，你将获益终身！",
		FileName:      "fileName.gif",
		FileId:        0,
	}
	fmt.Println(tools.ToJson(req))
}
