package main

import (
	"checkin-be/internal/controller"
	"checkin-be/pkg/invoker"

	"github.com/gin-gonic/gin"
)

func main() {
	invoker.Init()

	// Create a new gin router
	router := gin.Default()
	// Register the routes
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	api.GET("/checkin-users", controller.CheckinUsers)                              // 签到用户列表
	api.POST("/checkin", controller.Checkin)                                        // 签到
	api.DELETE("/checkin-users/:userid", controller.DeleteCheckinUser)              // 删除有问题的签到数据
	api.POST("/generate-exchange-card-pairs", controller.GenerateExchangeCardPairs) // 生成交换贺卡列表
	api.GET("/exchange-card-pairs", controller.ExchangeCardPairs)                   // 列出交换贺卡列表
	api.POST("/generate-lottery", controller.GenerateLottery)                       // 点击生成抽奖
	api.GET("/lotteries", controller.Lotteries)                                     // 抽奖结果
	api.POST("/lottery-setup", controller.LotterySetup)                             // 抽奖设置

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
