package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username"`
	Avatar   string `gorm:"avatar"`
	Poster   string `gorm:"poster"`
}

func (u *User) ToView() map[string]interface{} {
	return map[string]interface{}{
		"id":       u.ID,
		"username": u.Username,
		"avatar":   u.Avatar,
		"poster":   u.Poster,
	}
}

type ExchangeCardPair struct {
	RecvID uint `gorm:"recv_id"`
	SendID uint `gorm:"send_id"`
}

func (e *ExchangeCardPair) ToView() map[string]interface{} {
	return map[string]interface{}{
		"recv_id": e.RecvID,
		"send_id": e.SendID,
	}
}

type LotterySetup struct {
	gorm.Model
	Description string // 奖项描述
	Label       string // 奖项等级
	Amount      int32  // 奖项数量
}

type LotteryResult struct {
	Label  string `gorm:"label"`
	UserID uint   `gorm:"user_id"`
}

func (l *LotteryResult) ToView() map[string]interface{} {
	return map[string]interface{}{
		"label":   l.Label,
		"user_id": l.UserID,
	}
}
