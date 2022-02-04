package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var client sarama.SyncProducer
var msgChan chan *sarama.ProducerMessage

func Init(address []string, chanSize int64) (err error) {
	//1. 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	fmt.Println(address)

	//2. 连接 kafka
	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		logrus.Error("kafkaL:producer closed, err:", err)
		return
	}

	msgChan = make(chan *sarama.ProducerMessage, chanSize)
	go sendMsg()
	return
}

func sendMsg() {
	for {
		select {
		case msg := <-msgChan:
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				logrus.Warning("send msg failed,err:", err)
				return
			}
			logrus.Info("send msg to kafka success. pid:%v offset:", pid, offset)
		}
	}
}
