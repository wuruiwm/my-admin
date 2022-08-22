package model

import (
	"app/util"
	"gorm.io/gorm"
)

type Base struct {
	Id         string `gorm:"column:id;type:char(36);primary_key" json:"id"`
	CreateTime string `gorm:"column:create_time;type:datetime" json:"create_time"` // 创建时间
	UpdateTime string `gorm:"column:update_time;type:datetime" json:"update_time"` // 修改时间
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = util.Uuid()
	b.CreateTime = util.GetDate()
	b.UpdateTime = util.GetDate()
	return
}

func (b *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdateTime = util.GetDate()
	return
}
