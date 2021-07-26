package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
)

type OldCards struct {
	ID         int        `gorm:"column:id;primary_key"` //
	ProductID  int        `gorm:"column:product_id"`     //
	CardInfo   string     `gorm:"column:card_info"`      //
	CardStatus int        `gorm:"column:card_status"`    //
	CreatedAt  utils.Time `gorm:"column:created_at"`     //
	UpdatedAt  utils.Time `gorm:"column:updated_at"`     //
}

// TableName sets the insert table name for this struct type
func (o *OldCards) TableName() string {
	return "cards"
}

// GetMaxCardsID 获得最大卡密id
func GetMaxCardsID() (int, error) {
	var card OldCards
	err := utils.OldDB.Model(card).Order("id desc").First(&card).Error
	return card.ID, err
}

// GetCards 获得卡密
func GetCards(min, max int) ([]OldCards, error) {
	var cards []OldCards
	err := utils.OldDB.Model(cards).Where("id > ?", min).Where("id <= ?", max).Order("id ASC").Find(&cards).Error
	return cards, err
}
