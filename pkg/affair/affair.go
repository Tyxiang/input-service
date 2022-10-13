package affair

import (
	// "fmt"
	"errors"
	"time"
	// "math"

	"valve-test-device-software/pkg/serial"
)

type Config struct {
	intervalFor		time.Duration
	intervalAnswer	time.Duration
}

type Affair struct { 
	c 		Config
	cmd		chan []byte
	ans		chan []byte
	err		chan error
}

func NewAffair() *Affair {
	var a Affair
	a.cmd = make(chan []byte)
	a.ans = make(chan []byte)
	a.err = make(chan error)
	return &a
}

func (this *Affair) Setup(config map[string]interface{}) error {
	//
	intervalFor := 0
	ifor, ok := config["interval-for"]
	if ok {
		intervalFor = ifor.(int)
	}
	if intervalFor < 0 {
		err := errors.New("interval-for cannot be less than 0")
		return err
	}
	this.c.intervalFor = time.Duration(intervalFor) * time.Millisecond
	//
	intervalAnswer := 10
	ia, ok := config["interval-answer"]
	if ok {
		intervalAnswer = ia.(int)
	}
	if intervalAnswer < 0 {
		err := errors.New("interval-answer cannot be less than 0")
		return err
	}
	this.c.intervalAnswer = time.Duration(intervalAnswer) * time.Millisecond
	//
	go this.handle()
	//
	return nil
}
/* 
- 竭力模式；
- 通过 cmd 通道输入命令；
- 通过 ans 和 err 通道返回数据；
- 本身无 err 可能；
*/
func (this *Affair) handle() {
	for {
		time.Sleep(this.c.intervalFor)
		cmd := <- this.cmd
		err := serial.Write(cmd)
		if err != nil {
			this.err <- err
			serial.ReOpen()
			continue
		}
		this.err <- nil
		time.Sleep(this.c.intervalAnswer)
		ans, err := serial.Read() 
		if err != nil {
			this.err <- err
			serial.ReOpen()
			continue
		}
		this.err <- nil
		this.ans <- ans
	}
}

func (this *Affair) Send(cmd []byte) ([]byte, error) {
	// cmd
	this.cmd <- cmd
	// cmd err
	err := <- this.err
	if  err != nil {
		return nil, err
	}
	// ans err
	err = <- this.err 
	if err != nil {
		return nil, err
	}
	// ans
	ans := <- this.ans 
	return ans, nil
}
