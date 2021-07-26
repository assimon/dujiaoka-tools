package main

import (
	"github.com/assimon/dujiaoka_migrate/cmd"
	"github.com/assimon/dujiaoka_migrate/conf"
)

func main() {
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}
	cmd.InitCli()
}
