package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
)

type NewOrders struct {
	ID                     int        `gorm:"column:id;primary_key"`           //
	OrderSn                string     `gorm:"column:order_sn"`                 //订单号
	GoodsID                int        `gorm:"column:goods_id"`                 //关联商品id
	CouponID               int        `gorm:"column:coupon_id"`                //关联优惠码id
	Title                  string     `gorm:"column:title"`                    //订单名称
	Type                   int        `gorm:"column:type"`                     //1自动发货 2人工处理
	GoodsPrice             float64    `gorm:"column:goods_price"`              //商品单价
	BuyAmount              int        `gorm:"column:buy_amount"`               //购买数量
	CouponDiscountPrice    float64    `gorm:"column:coupon_discount_price"`    //优惠码优惠价格
	WholesaleDiscountPrice float64    `gorm:"column:wholesale_discount_price"` //批发价优惠
	TotalPrice             float64    `gorm:"column:total_price"`              //订单总价
	ActualPrice            float64    `gorm:"column:actual_price"`             //实际支付价格
	SearchPwd              string     `gorm:"column:search_pwd"`               //查询密码
	Email                  string     `gorm:"column:email"`                    //下单邮箱
	Info                   string     `gorm:"column:info"`                     //订单详情
	PayID                  int        `gorm:"column:pay_id"`                   //支付通道id
	BuyIP                  string     `gorm:"column:buy_ip"`                   //购买者下单IP地址
	TradeNo                string     `gorm:"column:trade_no"`                 //第三方支付订单号
	Status                 int        `gorm:"column:status"`                   //1待支付 2待处理 3处理中 4已完成 5处理失败 6异常 -1过期
	CouponRetBack          int        `gorm:"column:coupon_ret_back"`          //优惠码使用次数是否已经回退 0否 1是
	CreatedAt              utils.Time `gorm:"column:created_at"`               //
	UpdatedAt              utils.Time `gorm:"column:updated_at"`               //
	DeletedAt              utils.Time `gorm:"column:deleted_at"`               //
}

// TableName sets the insert table name for this struct type
func (n *NewOrders) TableName() string {
	return "orders"
}

// CreateOrders 新增订单
func CreateOrders(orders []NewOrders) error {
	err := utils.NewDB.Model(orders).Create(orders).Error
	return err
}
