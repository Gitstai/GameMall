package handler

import (
	"GameMall/config"
	"GameMall/logs"
	"GameMall/model"
	"GameMall/tools"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckLogin(c *gin.Context) {
	token := c.GetHeader("xx-token")
	if token == "" {
		logs.Logger.Info("cookie中无token")
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		c.Abort()
		return
	}
	id, isOk := tools.AuthCheck(token)
	if !isOk {
		logs.Logger.Info("token验证错误")
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		c.Abort()
		return
	}
	userInfo, err := model.GetTUserInfo(&model.TUser{Id: id})
	if err != nil {
		logs.Logger.Infof("系统错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		c.Abort()
		return
	} else if len(userInfo) != 1 {
		logs.Logger.Infof("找不到该用户, id:%v", id)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		c.Abort()
		return
	}
	logs.Logger.Infof("CheckLogin UserInfo:%v", tools.ToJson(userInfo))
	c.Set(config.UserInfo, userInfo[0])
	c.Next()
}

func GetUser(c *gin.Context) *model.TUser {
	if c == nil {
		return nil
	}
	userInfo, exists := c.Get(config.UserInfo)
	if !exists {
		return nil
	}
	realUserInfo, ok := userInfo.(*model.TUser)
	if !ok {
		return nil
	}
	return realUserInfo
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Header("Access-Control-Allow-Origin", "*")
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Token")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
