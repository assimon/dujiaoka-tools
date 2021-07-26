package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
)

type NewCarmis struct {
	ID        int        `gorm:"column:id;primary_key"` //
	GoodsID   int        `gorm:"column:goods_id"`       //所属商品
	Status    int        `gorm:"column:status"`         //状态 1未售出 2已售出
	Carmi     string     `gorm:"column:carmi"`          //卡密
	CreatedAt utils.Time `gorm:"column:created_at"`     //
	UpdatedAt utils.Time `gorm:"column:updated_at"`     //
	DeletedAt utils.Time `gorm:"column:deleted_at"`     //
}

// TableName sets the insert table name for this struct type
func (n *NewCarmis) TableName() string {
	return "carmis"
}

// CreateCarmis 新增卡密
func CreateCarmis(carmis []NewCarmis) error {
	err := utils.NewDB.Model(carmis).Create(carmis).Error
	return err
}
