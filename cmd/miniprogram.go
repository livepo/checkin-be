package main

import (
	"fmt"
	"time"

	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

type Cache struct{}

func (c Cache) Get(key string) interface{} {
	return key
}
func (c Cache) Set(key string, val interface{}, timeout time.Duration) error {
	return nil
}
func (c Cache) IsExist(key string) bool {
	return true
}

func (c Cache) Delete(key string) error {
	return nil
}

func main() {
	mini := miniprogram.NewMiniProgram(&config.Config{
		AppID:     "wxb1f8362f8273f030",
		AppSecret: "8a2412090a00af4994517388d38e246d",
		Cache:     Cache{},
	})
	auth := mini.GetAuth()
	resp, err := auth.Code2Session("061Mdr000lp6ON1B7K300Y7qVE0Mdr0w")
	fmt.Println(resp.UnionID, resp.OpenID, resp.SessionKey, err)
	fmt.Println(auth.GetPhoneNumber("061Mdr000lp6ON1B7K300Y7qVE0Mdr0w"))

}
