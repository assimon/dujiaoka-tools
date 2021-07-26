package handles

import (
	"github.com/assimon/dujiaoka_migrate/models"
	"github.com/gookit/color"
)

//ForMaxNum  每次迁移最大条数
const ForMaxNum = 200

func MigrateVersionToTwo() error {
	var err error
	/*color.Warn.Println("开始迁移商品分类...")
	err = migrateGoodsGroup()
	if err != nil {
		color.Error.Printf("商品分类迁移失败，已终止。 错误：%s", err.Error())
		return err
	}
	color.Success.Println("商品分类迁移成功！")
	color.Warn.Println("开始迁移商品...")
	err = migrateGoods()
	if err != nil {
		color.Error.Printf("商品迁移失败，已终止。 错误：%s", err.Error())
		return err
	}
	color.Success.Println("商品迁移成功！")
	color.Warn.Println("开始迁移优惠券...")
	err = migrateCoupons()
	if err != nil {
		color.Error.Printf("优惠券迁移失败，已终止。 错误：%s", err.Error())
		return err
	}
	color.Success.Println("优惠券迁移成功！")
	color.Warn.Println("开始迁移卡密...")
	err = migrateCarmis()
	if err != nil {
		color.Error.Printf("卡密迁移失败，已终止。 错误：%s", err.Error())
		return err
	}
	color.Success.Println("卡密迁移成功！")*/
	color.Warn.Println("开始迁移订单...")
	err = migrateOrders()
	if err != nil {
		color.Error.Printf("订单迁移失败，已终止。 错误：%s", err.Error())
		return err
	}
	color.Success.Println("订单迁移成功！")

	return err
}

// migrateGoodsGroup 迁移分类
func migrateGoodsGroup() error {
	oldClass, err := models.GetAllClassifys()
	if err != nil {
		return err
	}
	var goodsGroup []models.NewGoodsGroup
	for _, c := range oldClass {
		tempGroup := models.NewGoodsGroup{
			ID:        c.ID,
			GpName:    c.Name,
			IsOpen:    c.CStatus,
			Ord:       c.Ord,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		}
		goodsGroup = append(goodsGroup, tempGroup)
	}
	if len(goodsGroup) > 0 {
		err = models.CreateGroup(goodsGroup)
	}
	return err
}

// migrateGoods 迁移商品
func migrateGoods() error {
	products, err := models.GetAllProducts()
	if err != nil {
		return err
	}
	var goods []models.NewGoods
	for _, product := range products {
		tempGoods := models.NewGoods{
			ID:                product.ID,
			GroupID:           product.PdClass,
			GdName:            product.PdName,
			Picture:           product.PdPicture,
			RetailPrice:       product.CostPrice,
			ActualPrice:       product.ActualPrice,
			InStock:           product.InStock,
			SalesVolume:       product.SalesVolume,
			Ord:               product.Ord,
			BuyPrompt:         product.BuyPrompt,
			Description:       product.PdInfo,
			Type:              product.PdType,
			WholesalePriceCnf: product.WholesalePrice,
			OtherIpuCnf:       product.OtherIpu,
			IsOpen:            product.PdStatus,
			CreatedAt:         product.CreatedAt,
			UpdatedAt:         product.UpdatedAt,
			DeletedAt:         product.DeletedAt,
		}
		goods = append(goods, tempGoods)
	}
	if len(goods) > 0 {
		err = models.CreateGoods(goods)
	}
	return err
}

// migrateCoupons 迁移优惠券
func migrateCoupons() error {
	oldCoupons, err := models.GetAllCoupon()
	if err != nil {
		return err
	}
	var coupons []models.NewCoupons
	for _, coupon := range oldCoupons {
		tempCoupon := models.NewCoupons{
			ID:        coupon.ID,
			Discount:  coupon.Discount,
			IsUse:     coupon.IsStatus,
			IsOpen:    1,
			Coupon:    coupon.Card,
			Ret:       coupon.Ret,
			CreatedAt: coupon.CreatedAt,
			UpdatedAt: coupon.UpdatedAt,
		}
		coupons = append(coupons, tempCoupon)
	}
	if len(coupons) > 0 {
		err = models.CreateCoupons(coupons)
	}
	return err
}

// migrateCarmis 迁移卡密
func migrateCarmis() error {
	maxID, err := models.GetMaxCardsID()
	if err != nil {
		return err
	}
	for i := 0; i < maxID; i += ForMaxNum {
		cards, err := models.GetCards(i, i+ForMaxNum)
		if err != nil {
			return err
		}
		var carmis []models.NewCarmis
		for _, card := range cards {
			tempCarmi := models.NewCarmis{
				ID:        card.ID,
				GoodsID:   card.ProductID,
				Status:    card.CardStatus,
				Carmi:     card.CardInfo,
				CreatedAt: card.CreatedAt,
				UpdatedAt: card.UpdatedAt,
			}
			carmis = append(carmis, tempCarmi)
		}
		err = models.CreateCarmis(carmis)
		if err != nil {
			return err
		}
	}
	return err
}

// migrateOrders 迁移订单
func migrateOrders() error {
	maxID, err := models.GetMaxOrderID()
	if err != nil {
		return err
	}
	for i := 0; i < maxID; i += ForMaxNum {
		orders, err := models.GetOrders(i, i+ForMaxNum)
		if err != nil {
			return err
		}
		var newOrders []models.NewOrders
		for _, order := range orders {
			tempOrder := models.NewOrders{
				ID:            order.ID,
				OrderSn:       order.OrderID,
				GoodsID:       order.ProductID,
				CouponID:      order.CouponID,
				Title:         order.OrdTitle,
				Type:          order.OrdClass,
				GoodsPrice:    order.ProductPrice,
				BuyAmount:     order.BuyAmount,
				TotalPrice:    order.OrdPrice,
				ActualPrice:   order.OrdPrice,
				SearchPwd:     order.SearchPwd,
				Email:         order.Account,
				Info:          order.OrdInfo,
				PayID:         order.PayWay,
				BuyIP:         order.BuyIP,
				TradeNo:       order.PayOrd,
				Status:        order.OrdStatus + 1,
				CouponRetBack: 0,
				CreatedAt:     order.CreatedAt,
				UpdatedAt:     order.UpdatedAt,
				DeletedAt:     order.DeletedAt,
			}
			newOrders = append(newOrders, tempOrder)
		}
		err = models.CreateOrders(newOrders)
		if err != nil {
			return err
		}
	}
	return err
}
