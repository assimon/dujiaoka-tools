package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
)

type OldProducts struct {
	ID             int        `gorm:"column:id;primary_key"`  //
	PdName         string     `gorm:"column:pd_name"`         //商品名称
	CostPrice      float64    `gorm:"column:cost_price"`      //
	ActualPrice    float64    `gorm:"column:actual_price"`    //实际售价
	WholesalePrice string     `gorm:"column:wholesale_price"` //
	InStock        int        `gorm:"column:in_stock"`        //
	SalesVolume    int        `gorm:"column:sales_volume"`    //
	PdPicture      string     `gorm:"column:pd_picture"`      //商品图片
	Ord            int        `gorm:"column:ord"`             //排序
	BuyPrompt      string     `gorm:"column:buy_prompt"`      //购买提示
	PdInfo         string     `gorm:"column:pd_info"`         //商品详情
	PdType         int        `gorm:"column:pd_type"`         //1自动发卡 2代充
	OtherIpu       string     `gorm:"column:other_ipu"`       //
	PdStatus       int        `gorm:"column:pd_status"`       //
	PdClass        int        `gorm:"column:pd_class"`        //
	IsopenCoupon   int        `gorm:"column:isopen_coupon"`   //
	CreatedAt      utils.Time `gorm:"column:created_at"`      //
	UpdatedAt      utils.Time `gorm:"column:updated_at"`      //
	DeletedAt      utils.Time `gorm:"column:deleted_at"`      //
}

// TableName sets the insert table name for this struct type
func (o *OldProducts) TableName() string {
	return "products"
}

// GetAllProducts 获得所有老商品
func GetAllProducts() ([]OldProducts, error) {
	var products []OldProducts
	err := utils.OldDB.Model(products).Find(&products).Error
	return products, err
}
