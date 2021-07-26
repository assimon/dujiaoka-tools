package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
)

type NewGoods struct {
	ID                int        `gorm:"column:id;primary_key"`      //
	GroupID           int        `gorm:"column:group_id"`            //所属分类id
	GdName            string     `gorm:"column:gd_name"`             //商品名称
	Picture           string     `gorm:"column:picture"`             //商品图片
	RetailPrice       float64    `gorm:"column:retail_price"`        //零售价
	ActualPrice       float64    `gorm:"column:actual_price"`        //实际售价
	InStock           int        `gorm:"column:in_stock"`            //库存
	SalesVolume       int        `gorm:"column:sales_volume"`        //销量
	Ord               int        `gorm:"column:ord"`                 //排序权重 越大越靠前
	BuyLimitNum       int        `gorm:"column:buy_limit_num"`       //限制单次购买最大数量，0为不限制
	BuyPrompt         string     `gorm:"column:buy_prompt"`          //购买提示
	Description       string     `gorm:"column:description"`         //商品描述
	Type              int        `gorm:"column:type"`                //商品类型  1自动发货 2人工处理
	WholesalePriceCnf string     `gorm:"column:wholesale_price_cnf"` //批发价配置
	OtherIpuCnf       string     `gorm:"column:other_ipu_cnf"`       //其他输入框配置
	APIHook           string     `gorm:"column:api_hook"`            //回调事件
	IsOpen            int        `gorm:"column:is_open"`             //是否启用，1是 0否
	CreatedAt         utils.Time `gorm:"column:created_at"`          //
	UpdatedAt         utils.Time `gorm:"column:updated_at"`          //
	DeletedAt         utils.Time `gorm:"column:deleted_at"`          //
}

// TableName sets the insert table name for this struct type
func (n *NewGoods) TableName() string {
	return "goods"
}

// CreateGoods 新增商品
func CreateGoods(goods []NewGoods) error {
	err := utils.NewDB.Model(goods).Create(goods).Error
	return err
}
