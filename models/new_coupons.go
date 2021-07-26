package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
)

type NewCoupons struct {
	ID        int        `gorm:"column:id;primary_key"` //
	Discount  float64    `gorm:"column:discount"`       //优惠金额
	IsUse     int        `gorm:"column:is_use"`         //是否已经使用 1未使用 2已使用
	IsOpen    int        `gorm:"column:is_open"`        //是否启用 1是 0否
	Coupon    string     `gorm:"column:coupon"`         //优惠码
	Ret       int        `gorm:"column:ret"`            //剩余使用次数
	CreatedAt utils.Time `gorm:"column:created_at"`     //
	UpdatedAt utils.Time `gorm:"column:updated_at"`     //
	DeletedAt utils.Time `gorm:"column:deleted_at"`     //
}

// TableName sets the insert table name for this struct type
func (n *NewCoupons) TableName() string {
	return "coupons"
}

// CreateCoupons 新增优惠券
func CreateCoupons(coupons []NewCoupons) error {
	err := utils.NewDB.Model(coupons).Create(coupons).Error
	return err
}
