package router

import (
	"GameMall/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(handler.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/edu/mall/login", handler.Login)       //test ok
	r.POST("/edu/mall/register", handler.Register) //test ok

	authorized := r.Group("/")
	authorized.Use(handler.CheckLogin)
	eduRouter := authorized.Group("/edu/mall")
	{
		eduRouter.GET("/logout", handler.Logout)                         //可不写 前端直接去除header
		eduRouter.GET("/getUserInfo", handler.GetUserInfo)               //test ok
		eduRouter.GET("/searchEduProducts", handler.SearchEduProducts)   //test ok
		eduRouter.POST("/upsertEduProduct", handler.UpsertEduProduct)    //test ok
		eduRouter.GET("/getProductDetail", handler.GetProductDetail)     //test ok
		eduRouter.GET("/getProductEditInfo", handler.GetProductEditInfo) //test ok
		eduRouter.GET("/getPurchaseRecords", handler.GetPurchaseRecords) //test ok
		eduRouter.POST("/recharge", handler.Recharge)                    //test ok
		eduRouter.GET("/purchase", handler.Purchase)                     //test ok
		eduRouter.GET("/checkPurchased", handler.CheckPurchased)         //test ok
	}
	return r
}
