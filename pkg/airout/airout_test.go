package airout

import (
	"testing"
	"fmt"
	"time"

	"valve-test-device-software/pkg/serial"
	"valve-test-device-software/pkg/affair"
	"valve-test-device-software/pkg/modbus"
	"valve-test-device-software/pkg/aim"
	"valve-test-device-software/pkg/dio"
	"valve-test-device-software/pkg/vfd"
	"valve-test-device-software/pkg/sensors"
	"valve-test-device-software/pkg/prv"
	"valve-test-device-software/pkg/pump"
)

func Test(t *testing.T) {
	// serial
	fmt.Println("serial.Setup()", "err:", serial.Setup(map[string]interface{}{}))
	// affair
	fmt.Println("affair.Setup()", "err:", affair.Setup(map[string]interface{}{}))
	// modbus
	fmt.Println("modbus.Setup()", "err:", modbus.Setup(map[string]interface{}{}))
	// aim
	fmt.Println("aim.Setup()", "err:", aim.Setup(map[string]interface{}{}))	
	// dio
	fmt.Println("dio.Setup()", "err:", dio.Setup(map[string]interface{}{}))
	// vfd
	fmt.Println("vfd.Setup()", "err:", vfd.Setup(map[string]interface{}{}))
	// sensors
	fmt.Println("sensors.Setup()", "err:", sensors.Setup(map[string]interface{} {
		"dp1": map[string]interface{} {
			"channel": 4,
			"min": 0.0,
			"min-offset": 0.0,
			"max": 500.0,
			"max-offset": 0.0,
		},
		"dp2": map[string]interface{} {
			"channel": 3,
			"min": 0.0,
			"min-offset": 0.0,
			"max": 500.0,
			"max-offset": 0.0,
		},
		"p": map[string]interface{} {
			"channel": 2,
			"min": 0.0,
			"min-offset": 0.0,
			"max": 600.0,
			"max-offset": 0.0,
		},
		"f": map[string]interface{} {
			"channel": 1,
			"min": 0.0,
			"min-offset": 0.0,
			"max": 6.361,
			"max-offset": 0.0,
		},
		"t": map[string]interface{} {
			"channel": 0,
			"min": -10.0,
			"min-offset": 0.0,
			"max": 50.0,
			"max-offset": 0.0,
		},
	}))
	// prv
	fmt.Println("prv.Setup()", "err:", prv.Setup(map[string]interface{}{}))
	// pump
	fmt.Println("pump.Setup()", "err:", pump.Setup(map[string]interface{}{}))
	// airout
	fmt.Println("airout.Setup()", "err:", Setup(map[string]interface{}{}))
	//
	// fmt.Println("airout.Start()")
	// go Start()
	// for i := 0; i <= 10; i++ {
	// 	time.Sleep(time.Millisecond * 500)
	// 	ready, end, err := airout.GetState()
	// 	fmt.Println("airout.GetState()", "ready:", ready, "end:", end, "err:", err)
	// }
	// fmt.Println("airout.End()")
	// go End()
	// for i := 0; i <= 10; i++ {
	// 	time.Sleep(time.Millisecond * 500)
	// 	ready, end, err := GetState()
	// 	fmt.Println("airout.GetState()", "ready:", ready, "end:", end, "err:", err)
	// }
}