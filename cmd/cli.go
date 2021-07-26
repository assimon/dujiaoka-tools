package cmd

import (
	"github.com/assimon/dujiaoka_migrate/handles"
	"github.com/assimon/dujiaoka_migrate/utils"
	"github.com/gookit/color"
	"github.com/gookit/gcli/v3"
)

func InitCli() {
	app := gcli.NewApp()
	app.Version = "1.0.0"
	app.Desc = "【独角数卡】更新迁移命令工具,QQ群：568679748"
	app.Logo.Text = "     _        _ _             _         \n  __| |_   _ (_|_) __ _  ___ | | ____ _ \n / _` | | | || | |/ _` |/ _ \\| |/ / _` |\n| (_| | |_| || | | (_| | (_) |   < (_| |\n \\__,_|\\__,_|/ |_|\\__,_|\\___/|_|\\_\\__,_|\n           |__/                         "
	color.Printf("%s\n", color.WrapTag(app.Logo.Text, app.Logo.Style))
	app.Add(&gcli.Command{
		Name:    "migrate-version-to-two",
		Desc:    "迁移数据到2.x版本",
		Aliases: []string{"mvtt"},
		Func: func(c *gcli.Command, args []string) error {
			err := utils.InitDB()
			if err != nil {
				return err
			}
			err = handles.MigrateVersionToTwo()
			return err
		},
	})
	app.Run(nil)
}
