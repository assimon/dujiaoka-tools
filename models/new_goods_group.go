package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
	"time"
)

type NewGoodsGroup struct {
	ID        int       `gorm:"column:id;primary_key"` //
	GpName    string    `gorm:"column:gp_name"`        //分类名称
	IsOpen    int       `gorm:"column:is_open"`        //是否启用，1是 0否
	Ord       int       `gorm:"column:ord"`            //排序权重 越大越靠前
	CreatedAt time.Time `gorm:"column:created_at"`     //
	UpdatedAt time.Time `gorm:"column:updated_at"`     //
}

// TableName sets the insert table name for this struct type
func (n *NewGoodsGroup) TableName() string {
	return "goods_group"
}

// Creates 批量新增
func CreateGroup(goodsGroup []NewGoodsGroup) error {
	err := utils.NewDB.Model(goodsGroup).Create(goodsGroup).Error
	return err
}
