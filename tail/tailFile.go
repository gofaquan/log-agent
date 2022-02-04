package tail

import (
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

var TailObj *tail.Tail

func Init(filename string) (err error) {
	config := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:      true,
		MustExist:   false,
		Poll:        true,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	}

	TailObj, err = tail.TailFile(filename, config)
	if err != nil {
		logrus.Errorf("tail: create tailObj for path:%s failed, err:%v\n", filename, err)
		return
	}
	return
}
