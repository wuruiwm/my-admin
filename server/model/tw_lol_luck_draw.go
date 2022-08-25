package model

//TwLolLuckDraw 台服lol幸运抽奖表
type TwLolLuckDraw struct {
	Base
	Prize    string `gorm:"column:prize;type:char(50);comment:抽奖结果;NOT NULL" json:"prize"`
	Response string `gorm:"column:response;type:text;comment:抽奖原始结果;NOT NULL" json:"response"`
}

func (m *TwLolLuckDraw) TableName() string {
	return "tw_lol_luck_draw"
}
