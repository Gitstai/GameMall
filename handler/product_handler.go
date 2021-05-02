package handler

import (
	"GameMall/config"
	"GameMall/dto"
	"GameMall/logs"
	"GameMall/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func SearchEduProducts(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.SearchEduProductsReq)
	err := c.ShouldBindQuery(req)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	if req.PageNum < 1 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 5
	}

	prods, total, err := model.QueryTProduct(&model.TProduct{ProviderName: req.Provider, ProductType: int8(req.ProductType), Keywords: req.Keywords}, req.PageNum, req.PageSize)
	if err != nil {
		logs.Logger.Errorf("func:%v, err:%v", "model.QueryTProduct", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	res := make([]*dto.EduProduct, 0, len(prods))
	for _, p := range prods {
		tmp := new(dto.EduProduct)
		tmp.ProductId = p.ProviderId
		tmp.Provider = p.ProviderName
		tmp.ProductType = int32(p.ProductType)
		tmp.ProductDesc = p.DescText
		tmp.ProductName = p.Name
		tmp.ProductImg = p.DescImg
		tmp.CreatedTime = p.CreateTime.Unix()
		tmp.SaleVolume = p.SaleVolume
		tmp.Inventory = p.Inventory

		res = append(res, tmp)
	}

	DataHandlerWithTotal(c, res, total)
}

func UpsertEduProduct(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.UpsertEduProductReq)
	err := c.ShouldBindJSON(req)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	if req.Id != 0 { //编辑产品
		//先查询是否有该条记录，并进行providerId和userId的验证
		//TODO 暂时不查

		//更新产品信息
		if err := model.UpdateTProduct(&model.TProduct{
			Id:            req.Id,
			Name:          req.ProductName,
			ProductType:   int8(req.ProductType),
			Status:        int8(req.Status),
			Price:         req.Price,
			AfterSaleText: req.AfterSaleText,
			DescText:      req.ProductDesc,
			DescImg:       strings.Join(req.BannerImgs, ";"),
		}); err != nil {
			logs.Logger.Errorf("func:%v, err:%v", "model.UpdateTProduct", err)
			ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
			return
		}
	} else { //创建产品
		p, err := model.InsertTProduct(&model.TProduct{
			ProviderId:    user.Id,
			ProviderName:  user.Nickname,
			Name:          req.ProductName,
			ProductType:   int8(req.ProductType),
			Status:        int8(req.Status),
			Price:         req.Price,
			AfterSaleText: req.AfterSaleText,
			DescText:      req.ProductDesc,
			DescImg:       strings.Join(req.BannerImgs, ";"),
		})
		if err != nil {
			logs.Logger.Errorf("func:%v, err:%v", "model.InsertTProduct", err)
			ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
			return
		}
		req.Id = p.Id
	}

	if req.FileId != 0 { //编辑产品对应的文件
		//先查询是否有该条记录
		//TODO 暂时不查

		//更新文件信息
		if err := model.UpdateTProductFile(&model.TProductFile{
			Id:       req.FileId,
			FileType: int8(req.FileType),
			FileName: req.FileName,
			Url:      req.FileUrl,
		}); err != nil {
			logs.Logger.Errorf("func:%v, err:%v", "model.UpdateTProductFile", err)
			ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
			return
		}
	} else { //创建产品对应的文件
		if _, err := model.InsertTProductFile(&model.TProductFile{
			FileType:  int8(req.FileType),
			FileName:  req.FileName,
			Url:       req.FileUrl,
			ProductID: req.Id,
		}); err != nil {
			logs.Logger.Errorf("func:%v, err:%v", "model.InsertTProductFile", err)
			ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
			return
		}
	}

	DataHandler(c, nil)
}

func GetProductDetail(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.GetProductDetailReq)
	err := c.BindQuery(req)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	prods, err := model.GetTProduct(&model.TProduct{Id: req.ProductId})
	if err != nil {
		logs.Logger.Errorf("func:%v, err:%v", "model.GetTProduct", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}
	if len(prods) == 0 {
		logs.Logger.Infof("len(prods) == 0, productId:%v", req.ProductId)
		DataHandler(c, nil)
		return
	}

	res := dto.ProductDetail{
		Id:          prods[0].Id,
		ProductName: prods[0].Name,
		ProductType: int32(prods[0].ProductType),
		Price:       prods[0].Price,
		Status:      int32(prods[0].Status),
		ProductDesc: prods[0].DescText,
		BannerImgs:  strings.Split(prods[0].DescImg, ";"),
		Inventory:   prods[0].Inventory,
		SaleVolume:  prods[0].SaleVolume,
		Files:       make([]*dto.File, 0),
	}

	files, err := model.GetTProductFile(&model.TProductFile{
		ProductID: req.ProductId,
	})
	if err != nil {
		logs.Logger.Errorf("func:%v, err:%v", "model.GetTProduct", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	for _, file := range files {
		tmp := new(dto.File)
		tmp.FileId = file.Id
		tmp.FileUrl = file.Url
		tmp.FileName = file.FileName
		tmp.FileType = int32(file.FileType)

		res.Files = append(res.Files, tmp)
	}
	DataHandler(c, res)
	return
}

func GetProductEditInfo(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.GetProductEditInfoReq)
	err := c.BindQuery(req)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	prods, err := model.GetTProduct(&model.TProduct{Id: req.ProductId})
	if err != nil {
		logs.Logger.Errorf("func:%v, err:%v", "model.GetTProduct", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}
	if len(prods) == 0 {
		logs.Logger.Infof("len(prods) == 0, productId:%v", req.ProductId)
		DataHandler(c, nil)
		return
	}

	res := dto.ProductEditInfo{
		Id:            prods[0].Id,
		ProductName:   prods[0].Name,
		ProductType:   int32(prods[0].ProductType),
		Price:         prods[0].Price,
		Status:        int32(prods[0].Status),
		ProductDesc:   prods[0].DescText,
		BannerImgs:    strings.Split(prods[0].DescImg, ";"),
		Inventory:     prods[0].Inventory,
		Files:         make([]*dto.File, 0),
		AfterSaleText: prods[0].AfterSaleText,
	}

	files, err := model.GetTProductFile(&model.TProductFile{
		ProductID: req.ProductId,
	})
	if err != nil {
		logs.Logger.Errorf("func:%v, err:%v", "model.GetTProduct", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	for _, file := range files {
		tmp := new(dto.File)
		tmp.FileId = file.Id
		tmp.FileUrl = file.Url
		tmp.FileName = file.FileName
		tmp.FileType = int32(file.FileType)

		res.Files = append(res.Files, tmp)
	}
	DataHandler(c, res)
	return
}
