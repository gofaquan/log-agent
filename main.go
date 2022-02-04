package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/gofaquan/kafka"
	"github.com/gofaquan/tail"
	"github.com/sirupsen/logrus"
)

type config struct {
	KafkaConfig `ini:"kafka"`
	TaiLConfig  `ini:"tail"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	ChanSize int64  `ini:"chan_size"`
}

type TaiLConfig struct {
	Filepath string `ini:"file_path"`
	Topic    string `ini:"topic"`
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

	//2. 初始化 kafka 连接
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Error("init kafka failed, err: %v", err)
		return
	}
	fmt.Println("init kafka success !")

	//3. 根据配置中的日志路径初始化tail
	err = tail.Init(configObj.TaiLConfig.Filepath)
	if err != nil {
		logrus.Error("init tail failed, err:", err)
		return
	}
	fmt.Println("init tail success !")

	err = kafka.Run(configObj.TaiLConfig.Topic)
	if err != nil {
		logrus.Error("获取 topic 出错:", err)
		return
	}
}
