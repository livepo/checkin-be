package controller

import (
	"checkin-be/internal/model"
	"checkin-be/pkg/invoker"
	"math/rand"
	"time"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

type CheckinRequest struct {
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
	Poster   string `json:"poster"`
}

// Checkin 用户签到
func Checkin(c *gin.Context) {
	var req CheckinRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "req params error",
		})
		return
	}
	user := model.User{}
	if err := invoker.DB.Model(model.User{}).First(&user, "username = ?", req.UserName).Error; err != nil {
		user.UserName = req.UserName
		user.Avatar = req.Avatar
		user.Poster = req.Poster
	}
	if err := invoker.DB.Create(&user).Error; err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "db error",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

// CheckinUsers 签到信息列表
func CheckinUsers(c *gin.Context) {
	users := make([]*model.User, 0)
	if err := invoker.DB.Model(model.User{}).Find(&users).Error; err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "db error",
		})
		return
	}
	resp := make([]map[string]interface{}, 0)
	for _, u := range users {
		resp = append(resp, u.ToView())
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": resp,
	})
}

// GenerateExchangeCardPairs 根据现有签到人数生成交换贺卡列表
func GenerateExchangeCardPairs(c *gin.Context) {
	invoker.DB.Where("1 = 1").Delete(&model.ExchangeCardPair{})
	users := make([]*model.User, 0)
	if err := invoker.DB.Model(model.User{}).Find(&users).Error; err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "db error",
		})
		return
	}
	arr := make([]int, 0)
	for i := len(users); i > 0; i-- {
		arr = append(arr, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	pairs := make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		pair := make([]int, 2)
		pair[0] = arr[i]
		if i == len(arr)-1 {
			pair[1] = arr[0]
		} else {
			pair[1] = arr[i+1]
		}
		pairs = append(pairs, pair)
	}
	for _, pair := range pairs {
		ecp := model.ExchangeCardPair{
			RecvID: cast.ToUint(pair[0]),
			SendID: cast.ToUint(pair[1]),
		}
		invoker.DB.Create(ecp)
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

// ExchangeCardPairs 列出交换贺卡列表
func ExchangeCardPairs(c *gin.Context) {
	ecps := make([]*model.ExchangeCardPair, 0)
	if err := invoker.DB.Model(model.ExchangeCardPair{}).Find(&ecps).Error; err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "db error",
		})
		return
	}
	resp := make([]map[string]interface{}, 0)
	for _, e := range ecps {
		resp = append(resp, e.ToView())
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": resp,
	})
}

type GenerateLotteryRequest struct {
	Label string
}

// GenerateLottery 生成抽奖结果
func GenerateLottery(c *gin.Context) {
	var req GenerateLotteryRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "req params error",
		})
		return
	}
	ls := model.LotterySetup{}
	if err := invoker.DB.Model(model.LotterySetup{}).Last(&ls, "label = ?", req.Label).Error; err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "error, setup lottery label first",
		})
		return
	}
	invoker.DB.Where("label = ?", req.Label).Delete(&model.LotteryResult{})
	// 从未中过奖的人中选出`ls.Amount`数量的人来中奖
	users := make([]*model.User, 0)
	invoker.DB.Model(model.User{}).Where("id not in (?)", invoker.DB.Model(model.LotteryResult{}).Select("user_id")).Find(&users)
	if len(users) < cast.ToInt(ls.Amount) {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "number of users who never be choose has not enough, setup lottery label and amount again",
		})
		return
	}
	arr := make([]int, 0)
	for i := len(users); i > 0; i-- {
		arr = append(arr, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	chosen := make([]*model.User, 0)
	for _, i := range arr[:ls.Amount] {
		chosen = append(chosen, users[i])
	}
	for _, c := range chosen {
		lr := model.LotteryResult{
			Label:  ls.Label,
			UserID: c.ID,
		}
		invoker.DB.Create(&lr)
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

type LotterySetupRequest struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Amount      int32  `json:"amount"`
}

// LotterySetup 抽奖设置
func LotterySetup(c *gin.Context) {
	var req LotterySetupRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "req params error",
		})
		return
	}
	ls := model.LotterySetup{
		Label:       req.Label,
		Description: req.Description,
		Amount:      req.Amount,
	}
	if err := invoker.DB.Create(&ls).Error; err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "create lottery setup error",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

// Lotteries 查询抽奖结果
func Lotteries(c *gin.Context) {
	label := c.Query("label")
	if label == "" {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "params error",
		})
		return
	}

	results := make([]*model.LotteryResult, 0)
	invoker.DB.Model(model.LotteryResult{}).Where("label = ?", label).Find(&results)
	resp := make([]map[string]interface{}, 0)
	for _, r := range results {
		resp = append(resp, r.ToView())
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": resp,
	})
}