package utils

import (
	"github.com/assimon/dujiaoka_migrate/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	NewDB *gorm.DB
	OldDB *gorm.DB
)

func getConn(dbDns string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbDns), &gorm.Config{})
	return db, err
}

func InitDB() error {
	var err error
	NewDB, err = getConn(conf.NewDBDns)
	OldDB, err = getConn(conf.OldDBDns)
	return err
}
