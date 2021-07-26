package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
)

type OldCoupons struct {
	ID        int        `gorm:"column:id;primary_key"` //
	ProductID int        `gorm:"column:product_id"`     //所属商品
	CType     int        `gorm:"column:c_type"`         //1为一次性使用 2为重复使用
	Discount  float64    `gorm:"column:discount"`       //优惠金额
	IsStatus  int        `gorm:"column:is_status"`      //是否已经使用1正常  2已使用
	Card      string     `gorm:"column:card"`           //优惠券内容
	Ret       int        `gorm:"column:ret"`            //剩余可用次数
	CreatedAt utils.Time `gorm:"column:created_at"`     //
	UpdatedAt utils.Time `gorm:"column:updated_at"`     //
}

// TableName sets the insert table name for this struct type
func (o *OldCoupons) TableName() string {
	return "coupons"
}

// GetAllCoupon 获得所有优惠券
func GetAllCoupon() ([]OldCoupons, error) {
	var coupons []OldCoupons
	err := utils.OldDB.Model(coupons).Find(&coupons).Error
	return coupons, err
}
