package aim

import (
	"testing"
	"fmt"
	"time"

	"valve-test-device-software/pkg/serial"
	"valve-test-device-software/pkg/affair"
	"valve-test-device-software/pkg/modbus"
)

func Test(t *testing.T) {
	// serial
	fmt.Println("serial.Setup()", "err:", serial.Setup(map[string]interface{} {
		// "name": "COM3",
		// "baud": 9600,
		// "size": 8,
		// "parity": "N",
		// "stopbits": 1,
		// "timeout": 1000,
	}))
	// affair
	fmt.Println("affair.Setup()", "err:", affair.Setup(map[string]interface{} {}))
	// modbus
	fmt.Println("modbus.Setup()", "err:", modbus.Setup(map[string]interface{} {}))
	// aim
	fmt.Println("aim.Setup()", "err:", Setup(map[string]interface{} {
		// "address": 1,
		// "interval-for": 500,
	}))
	//
	time.Sleep(time.Millisecond * 1000)
	//
	for i := 0; i < 8; i++ {
		time.Sleep(time.Millisecond * 1000)
		s, err := GetSignal(i)
		fmt.Println("aim.GetSignal", "s:", s, "err:", err)
	}
}
