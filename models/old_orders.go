package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
)

type OldOrders struct {
	ID           int        `gorm:"column:id;primary_key"` //
	OrderID      string     `gorm:"column:order_id"`       //订单号
	ProductID    int        `gorm:"column:product_id"`     //关联所属商品
	CouponID     int        `gorm:"column:coupon_id"`      //优惠券id
	OrdClass     int        `gorm:"column:ord_class"`      //1自动发卡 2代充
	ProductPrice float64    `gorm:"column:product_price"`  //商品单价
	OrdPrice     float64    `gorm:"column:ord_price"`      //
	BuyAmount    int        `gorm:"column:buy_amount"`     //购买数量
	OrdTitle     string     `gorm:"column:ord_title"`      //订单名称
	SearchPwd    string     `gorm:"column:search_pwd"`     //查询密码
	Account      string     `gorm:"column:account"`        //充值账号
	OrdInfo      string     `gorm:"column:ord_info"`       //订单详情
	PayOrd       string     `gorm:"column:pay_ord"`        //第三发支付id
	PayWay       int        `gorm:"column:pay_way"`        //第三方支付方式
	BuyIP        string     `gorm:"column:buy_ip"`         //购买者ip地址
	OrdStatus    int        `gorm:"column:ord_status"`     //1待处理 2已处理 3已完成  4处理失败
	CreatedAt    utils.Time `gorm:"column:created_at"`     //
	UpdatedAt    utils.Time `gorm:"column:updated_at"`     //
	DeletedAt    utils.Time `gorm:"column:deleted_at"`     //
}

// TableName sets the insert table name for this struct type
func (o *OldOrders) TableName() string {
	return "orders"
}

// GetMaxOrderID 获得最大订单id
func GetMaxOrderID() (int, error) {
	var order OldOrders
	err := utils.OldDB.Model(order).Order("id desc").First(&order).Error
	return order.ID, err
}

// GetOrders 获得订单
func GetOrders(min, max int) ([]OldOrders, error) {
	var orders []OldOrders
	err := utils.OldDB.Model(orders).Where("id > ?", min).Where("id <= ?", max).Order("id ASC").Find(&orders).Error
	return orders, err
}
