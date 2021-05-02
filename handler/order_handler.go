package handler

import (
	"GameMall/config"
	"GameMall/dto"
	"GameMall/logs"
	"GameMall/model"
	"GameMall/service"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetPurchaseRecords(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.GetPurchaseRecordsReq)
	err := c.BindQuery(req)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	//if req.UserId != user.Id {
	//	logs.Logger.Infof("req.UserId:%v, user.Id:%v, not equal", req.UserId, user.Id)
	//	ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
	//	return
	//}
	if req.PageNum < 1 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 5
	}

	records, count, err := model.QueryTPurchaseRecords(&model.TPurchaseRecords{UserId: user.Id, ProductName: req.ProductName}, req.PageNum, req.PageSize)
	if err != nil {
		logs.Logger.Errorf("func:%v, err:%v", "model.GetTPurchaseRecords", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	res := make([]*dto.PurchaseRecord, 0, len(records))
	for _, record := range records {
		tmp := new(dto.PurchaseRecord)
		tmp.RecordId = record.Id
		tmp.ProductId = record.ProductId
		tmp.ProductName = record.ProductName
		tmp.CreatedTime = record.CreateTime.Unix()
		tmp.ProductDesc = record.ProductDesc
		tmp.DescImg = strings.Split(record.DescImg, ";")
		tmp.ProductType = int32(record.ProductType)
		tmp.Payment = record.Payment

		res = append(res, tmp)
	}

	DataHandlerWithTotal(c, res, count)
	return
}

func Recharge(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.RechargeReq)
	err := c.ShouldBindJSON(req)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	////验证传过来的userid是否和登录的userid一致
	//if req.UserId == 0 || req.UserId != user.Id {
	//	logs.Logger.Infof("userId is not equal, req.UserId:%v, userId:%v", req.UserId, user.Id)
	//	ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
	//	return
	//}

	//验证激活码并获取其相应的金额
	amount := service.GetKeyAmount(req.CDkey)
	if amount <= 0 {
		logs.Logger.Errorf("fun:=service.GetKeyAmount, amount:%v, cdKey:%v", err, req.CDkey)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCdKeyNotLegal)
		return
	}

	//增加账户余额
	err = model.IncreaseBalance(user.Id, amount)
	if err != nil {
		logs.Logger.Errorf("fun:=model.IncreaseBalance, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	//有效增加账户余额后，使该激活码失效
	err = model.DelTCdKey(req.CDkey)
	if err != nil {
		logs.Logger.Errorf("fun:=model.DelTCdKey, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	//写入充值记录
	_, err = model.InsertTPrepaidRecords(&model.TPrepaidRecords{
		UserId: user.Id,
		CdKey:  req.CDkey,
		Amount: amount,
	})
	if err != nil {
		logs.Logger.Errorf("fun:=model.InsertTPrepaidRecords, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	DataHandler(c, nil)
	return
}

func Purchase(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.PurchaseReq)
	err := c.BindQuery(req)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	////验证传过来的userid是否和登录的userid一致
	//if req.UserId == 0 || req.UserId != user.Id {
	//	logs.Logger.Infof("userId is not equal, req.UserId:%v, userId:%v", req.UserId, user.Id)
	//	ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
	//	return
	//}

	//找到该产品，提取信息
	products, err := model.GetTProduct(&model.TProduct{Id: req.ProductId})
	if err != nil {
		logs.Logger.Infof("func:=model.GetTProduct, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}
	if len(products) == 0 {
		logs.Logger.Infof("func:=model.GetTProduct, len(products):%v", len(products))
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgNotExistProducts)
		return
	}

	product := products[0]

	//获取账户信息，余额对比产品价格
	if user.Balance < product.Price {
		logs.Logger.Infof("账户余额不足, user.Balance:%v, product.Price:%v", user.Balance, product.Price)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgLackOfBalance)
		return
	}

	//检查之前是否购买了
	records, err := model.GetTPurchaseRecords(&model.TPurchaseRecords{UserId: user.Id, ProductId: req.ProductId})
	if err != nil {
		logs.Logger.Errorf("func:=model.GetTPurchaseRecords, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	//之前购买过
	if len(records) > 0 {
		logs.Logger.Infof("func:=model.GetTPurchaseRecords, len(records):%v", len(records))
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgHaveBought)
		return
	}

	//写入购买记录
	if _, err = model.InsertTPurchaseRecords(&model.TPurchaseRecords{
		UserId:        user.Id,
		ProductId:     req.ProductId,
		Payment:       product.Price,
		ProductDesc:   product.DescText,
		ProductName:   product.Name,
		ProductType:   product.ProductType,
		DescImg:       product.DescImg,
		AfterSaleText: product.AfterSaleText,
	}); err != nil {
		logs.Logger.Errorf("func:=model.InsertTPurchaseRecords, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	//扣除余额
	if err = model.DecreaseBalance(user.Id, product.Price); err != nil {
		logs.Logger.Errorf("func:=model.DecreaseBalance, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	DataHandler(c, nil)
	return
}

func CheckPurchased(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.CheckPurchasedReq)
	err := c.BindQuery(req)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	////验证传过来的userid是否和登录的userid一致
	//if req.UserId == 0 || req.UserId != user.Id {
	//	logs.Logger.Infof("userId is not equal, req.UserId:%v, userId:%v", req.UserId, user.Id)
	//	ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
	//	return
	//}

	//查询购买记录
	records, err := model.GetTPurchaseRecords(&model.TPurchaseRecords{UserId: user.Id, ProductId: req.ProductId})
	if err != nil {
		logs.Logger.Errorf("func:=model.GetTPurchaseRecords, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	//无购买记录直接返回空
	if len(records) == 0 {
		DataHandler(c, nil)
		return
	}

	//有购买记录 则包装返回数据
	res := new(dto.AfterSaleInfo)
	files, err := model.GetTProductFile(&model.TProductFile{ProductID: req.ProductId})
	if err != nil {
		logs.Logger.Errorf("func:=model.GetTProductFile, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgBusinessException)
		return
	}

	res.ProductName = records[0].ProductName
	res.AfterSaleText = records[0].AfterSaleText
	res.Files = make([]*dto.File, 0, len(files))
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
