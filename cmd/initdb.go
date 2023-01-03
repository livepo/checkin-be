package main

import (
	"checkin-be/internal/model"
	"checkin-be/pkg/invoker"
)

func main() {
	invoker.Init()
	invoker.DB.AutoMigrate(&model.Movie{}, &model.User{}, &model.Node{})
}
