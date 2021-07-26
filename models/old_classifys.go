package models

import (
	"github.com/assimon/dujiaoka_migrate/utils"
	"time"
)

type OldClassifys struct {
	ID        int       `gorm:"column:id;primary_key"` //
	Name      string    `gorm:"column:name"`           //
	Icon      string    `gorm:"column:icon"`           //
	Ord       int       `gorm:"column:ord"`            //
	CStatus   int       `gorm:"column:c_status"`       //
	CreatedAt time.Time `gorm:"column:created_at"`     //
	UpdatedAt time.Time `gorm:"column:updated_at"`     //
}

// TableName sets the insert table name for this struct type
func (o *OldClassifys) TableName() string {
	return "classifys"
}

// GetAllClassifys 获得所有老分类
func GetAllClassifys() ([]OldClassifys, error) {
	var classifys []OldClassifys
	err := utils.OldDB.Model(classifys).Find(&classifys).Error
	return classifys, err
}
