package airout

import (
	// "fmt"
	"errors"
	"time"
	// "math"

	"valve-test-device-software/pkg/sensors"
	"valve-test-device-software/pkg/prv"
	"valve-test-device-software/pkg/pump"
)

type Config struct {
	intervalFor		time.Duration
	pumpMax			int
	airOutPressure	float64
	pressureReady 	float64
}

type Airout struct { 
	c		Config
	exit 	chan struct{}
	err 	chan error
}

func NewAirout() *Airout {
	var a Airout
	return &a
}

func (this *Airout) Setup(config map[string]interface{}) error {
	//
	intervalFor := 1000
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
	pumpMax := 20
	pm, ok := config["pump-max"]
	if ok {
		pumpMax = pm.(int)
	}
	if pumpMax < 0 {
		err := errors.New("pump-max cannot be less than 0")
		return err
	}
	this.c.pumpMax = pumpMax
	// 
	airOutPressure := 10.0
	aop, ok := config["air-out-pressure"]
	if ok {
		airOutPressure = aop.(float64)
	}
	if airOutPressure < 0 {
		err := errors.New("air-out-pressure cannot be less than 0")
		return err
	}
	this.c.airOutPressure = airOutPressure
	//
	pressureReady := 50.0
	pr, ok := config["pressure-ready"]
	if ok {
		pressureReady = pr.(float64)
	}
	if pressureReady < 0 {
		err := errors.New("pressure-ready cannot be less than 0")
		return err
	}
	this.c.pressureReady = pressureReady
	//
	return nil
}

func (this *Airout) Start() {
	// 准备
	this.err = make(chan error)
	this.exit = make(chan struct{})
	//// 停止水泵
	err := pump.Stop()
	if err != nil {
		this.err <- err
		return
	}
	err = pump.Speed(0)
	if err != nil {
		this.err <- err
		return
	}
	//// 关闭保压阀
	err = prv.Close()
	if err != nil {
		this.err <- err
		return
	}
	// 开始 
	fp := 0 	// 水泵速度
	pump.Start()
	loop:
	for {
		select {
		case <- this.exit:
			// 关闭水泵
			err := pump.Stop()
			if err != nil {
				this.err <- err
				return
			}
			err = pump.Speed(0)
			if err != nil {
				this.err <- err
				return
			}
			// 开启保压阀
			err = prv.Open()
			if err != nil {
				this.err <- err
				return
			}
			break loop
		default: 
			// 延时
			time.Sleep(this.c.intervalFor)
			// 获取实时压力
			p, err := sensors.GetValueP()
			if err != nil {
				this.err <- err
				return
			}
			// 未达到最低压力时，增加水泵转速
			if p <= this.c.pressureReady {
				fp++
				err := pump.Speed(fp)
				if err != nil {
					this.err <- err
					return
				}
			}
			// 达到排气压力时，开启保压阀
			if p > this.c.airOutPressure {
				err := prv.TurnOpen(5)
				if err != nil {
					this.err <- err
					return
				}
			}
			// 达到就绪压力时返回
			if p > this.c.pressureReady {
				break loop
			}
		}
	}
	this.err <- nil
	close(this.err)
	return
}

func (this *Airout) Exit() chan struct{} {
	return this.exit
}
func (this *Airout) Err() chan error {
	return this.err
}