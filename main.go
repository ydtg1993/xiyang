package main

import (
	"github.com/ydtg1993/ant"
	"pow/controller"
	"pow/robot"
	"pow/tools/config"
	"pow/tools/database"
	"pow/tools/log"
	"pow/tools/rd"
	"time"
)

func main() {
	Setup()
	//controller.MenuScan()
	controller.DetailScan()
	t := time.NewTicker(time.Second * 3000)
	<-t.C
}

func Setup() {
	err := config.Spe.SetUp()
	if err != nil {
		panic(err)
	}

	config.Spe.RedisDb = config.Spe.SourceId

	mylog := new(log.LogsManage)
	err = mylog.SetUp()
	if err != nil {
		panic(err)
	}

	db := new(database.MysqlManage)
	err = db.Setup()
	if err != nil {
		panic(err)
	}

	redisManage := new(rd.RedisManage)
	err = redisManage.SetUp()
	if err != nil {
		panic(err)
	}

	go ant.Build(config.Spe.Maxthreads, robot.GetSeleniumArgs())
}
