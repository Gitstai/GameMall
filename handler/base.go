package handler

import (
	"GameMall/config"
	"GameMall/dto"
	"GameMall/logs"
	"GameMall/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandler 基本错误
func ErrorHandler(c *gin.Context, st config.ErrorCode, msg string) {
	logs.Logger.Infof("resp no ok st:%v ,msg:%v ", st, msg)
	c.JSON(http.StatusOK, dto.NewResponseWithStatusData(st, msg, dto.EmptyObj{}))
}

// ErrorParamValidHandler 参数错误
func ErrorParamValidHandler(c *gin.Context, err error) {
	logs.Logger.Infof("resp param valid err:%v", err)
	c.JSON(http.StatusOK, dto.NewResponseWithStatusData(config.ErrCodeErrREQParamInvalid, "请求参数错误", dto.EmptyObj{}))
}

// ErrorHandlerData 需要返回数据的错误
func ErrorHandlerData(c *gin.Context, st config.ErrorCode, msg string, data interface{}) {
	logs.Logger.Infof("resp no ok code:%v ,msg:%v :%v", st, msg, tools.ToJson(data))
	c.JSON(http.StatusOK, dto.NewResponseWithStatusData(st, msg, data))
}

// DataHandler 返回正常数据
func DataHandler(c *gin.Context, data interface{}) {
	logs.Logger.Infof("resp ok :%v", tools.ToJson(data))
	c.JSON(http.StatusOK, dto.NewResponse(data))
}

// DataHandlerWithTotal 返回分页数据
func DataHandlerWithTotal(c *gin.Context, data interface{}, total int64) {
	logs.Logger.Infof("resp ok :%v", tools.ToJson(data))
	c.JSON(http.StatusOK, dto.NewResponseWithTotal(data, total))
}
