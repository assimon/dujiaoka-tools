package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"strings"
)

const (
	FilePath = "/conf/config.ini"
)

var (
	NewDBDns string
	OldDBDns string
)

func InitConfig() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	confPath := strings.Replace(pwd+FilePath, string(filepath.Separator), "/", -1)
	cfg, err := ini.Load(confPath)
	if err != nil {
		return err
	}
	newDBSect := cfg.Section("new_db")
	oldDBSect := cfg.Section("old_db")
	NewDBDns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		newDBSect.Key("user"),
		newDBSect.Key("password"),
		newDBSect.Key("host"),
		newDBSect.Key("port"),
		newDBSect.Key("database"),
	)
	OldDBDns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		oldDBSect.Key("user"),
		oldDBSect.Key("password"),
		oldDBSect.Key("host"),
		oldDBSect.Key("port"),
		oldDBSect.Key("database"),
	)
	return err
}
