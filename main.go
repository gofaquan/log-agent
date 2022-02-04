package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type config struct {
}

func main() {
	// 1. 读配置文件 `go-ini`
	var configObj = new(config)
	err := ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		logrus.Error("load config failed:%v", err)
	}
	fmt.Println(".ini file load success !")

}
