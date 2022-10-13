package aim

import (
	// "fmt"
	"errors"
	"time"
	"math"

	"valve-test-device-software/pkg/modbus"
)

type Config struct {
	address 		int
	intervalFor		time.Duration
}

type Aim struct { 
	c 			Config
	signals		[8]float64
	e			error
}

func NewAim() *Aim {
	var a Aim
	return &a
}

func (this *Aim) Setup(config map[string]interface{}) error {
	//
	address := 1
	a, ok := config["address"]
	if ok {
		address = a.(int)
	}
	if address < 0 {
		err := errors.New("address cannot be less than 0")
		return err
	}
	this.c.address = address
	//
	intervalFor := 100
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
	go this.Start()
	//
	return nil
}

/* 
- 竭力模式；
- 有 err 时，只保存不结束；
- 再有 err 时，覆盖；
- 再成功时，清空 err；
- 读取数据等操作时返回 err；
*/
func (this *Aim) Start() {
	for {
		// 延时
		time.Sleep(this.c.intervalFor)
		data, err := modbus.ReadMultipleHoldingRegisters(this.c.address, 40001, 8)
		if err != nil {
			this.e = err
			continue
		}
		var ss [8]float64
		ss[0] = inter(4, 20, 0, int(data[ 0])*256 + int(data[ 1]), 65535)
		ss[1] = inter(4, 20, 0, int(data[ 2])*256 + int(data[ 3]), 65535)
		ss[2] = inter(4, 20, 0, int(data[ 4])*256 + int(data[ 5]), 65535)
		ss[3] = inter(4, 20, 0, int(data[ 6])*256 + int(data[ 7]), 65535)
		ss[4] = inter(4, 20, 0, int(data[ 8])*256 + int(data[ 9]), 65535)
		ss[5] = inter(4, 20, 0, int(data[10])*256 + int(data[11]), 65535)
		ss[6] = inter(4, 20, 0, int(data[12])*256 + int(data[13]), 65535)
		ss[7] = inter(4, 20, 0, int(data[14])*256 + int(data[15]), 65535)
		this.signals = ss
		this.e = nil
	}
}

func (this *Aim) GetSignals() ([8]float64, error) {
	if this.e != nil {
		return [8]float64{}, this.e
	}
	return this.signals, nil
}
func (this *Aim) GetSignal(channel int) (float64, error) {
	if this.e != nil {
		return 0.0, this.e
	}
	return this.signals[channel], nil
}

func inter(a float64, c float64, x int, y int, z int) (b float64) {
	b = (a * (float64(z) - float64(y)) + c * (float64(y) - float64(x)))/(float64(z) - float64(x))
	b = math.Round(b*1000)/1000
	return
}
