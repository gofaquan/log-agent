package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/gofaquan/tail"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func Run(topic string) (err error) {
	for {
		line, ok := <-tail.TailObj.Lines //chan Lines into file
		if !ok {
			logrus.Warn("tail file 重新打开中, 路径为:%s\n", tail.TailObj.Filename)
			time.Sleep(time.Second)
			continue
		}
		// 如果是空行就略过, fmt.Printf("%#v\n", line.Text)输出为 \r ,故下面是 \r
		if len(strings.Trim(line.Text, "\r")) == 0 {
			logrus.Info("出现空行,直接跳过...")
			continue
		}

		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(line.Text),
		}
		//ToMsgChan(msg)
		msgChan <- msg
	}
}

//func ToMsgChan(message *sarama.ProducerMessage) {
//	msgChan <- message
//}
