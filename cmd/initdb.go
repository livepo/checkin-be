package main

import (
	"checkin-be/internal/model"
	"checkin-be/pkg/invoker"
)

func main() {
	invoker.Init()
	invoker.DB.AutoMigrate(&model.User{}, &model.ExchangeCardPair{}, &model.LotterySetup{}, &model.LotteryResult{})
}
