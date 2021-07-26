package main

import (
	"github.com/assimon/dujiaoka_migrate/cmd"
	"github.com/assimon/dujiaoka_migrate/conf"
	"github.com/assimon/dujiaoka_migrate/utils"
)

func main() {
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}
	err = utils.InitDB()
	if err != nil {
		panic(err)
	}
	cmd.InitCli()
}
