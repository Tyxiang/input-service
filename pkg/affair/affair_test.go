package affair

import (
	"testing"
	"fmt"
	"time"

	"valve-test-device-software/pkg/serial"
)

func Test(t *testing.T) {
	// serial
	fmt.Println("serial.Setup()", "err:", serial.Setup(map[string]interface{}{}))
	// affair
	fmt.Println("affair.Setup()", "err:", Setup(map[string]interface{}{
		// "interval-for": 0,
		// "interval-answer": 50,
	}))
	//
	go func() {
		for i := 0; i <= 5; i++ {
			ans, err := Send([]byte{0x02,0x02,0x00,0x00,0x00,0x01,0xB9,0xF9})
			fmt.Println("affair.Send1()", "ans:", ans, "err:", err)
		}
	}()
	go func() {
		for i := 0; i <= 5; i++ {
			ans, err := Send([]byte{0x02,0x02,0x00,0x00,0x00,0x01,0xB9,0xF9})
			fmt.Println("affair.Send2()", "ans:", ans, "err:", err)
		}
	}()
	//
	time.Sleep(time.Second * 2)
	//
}