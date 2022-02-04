package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/gofaquan/kafka"
	"github.com/sirupsen/logrus"
)

type config struct {
	KafkaConfig `ini:"kafka"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	ChanSize int64  `ini:"chan_size"`
}

func main() {
	// 1. 读配置文件 `go-ini`
	var configObj = new(config)
	err := ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		logrus.Error("load config failed:%v", err)
		return
	}
	fmt.Println(".ini file load success !")

	fmt.Println(configObj.KafkaConfig.Address)

	//2. 初始化 kafka 连接
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Error("init kafka failed, err: %v", err)
		return
	}
	fmt.Println("init kafka success !")

}
