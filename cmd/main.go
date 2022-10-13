package main

import (
	"fmt"
	"io"
	"os"
	"encoding/json"
	"net/url"
	"time"
	"embed"
	"net/http"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/go-vgo/robotgo"
	
	// "valve-test-device-software/pkg/aim"
)

func main() {
	// # init config
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("configs/")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		// panic(err)
		fmt.Println(err)
		os.Exit(1)
	}

	// # system log
	logDir := config.GetString("log.dir")
	_, err = os.Stat(logDir)
	if os.IsNotExist(err) {
		os.MkdirAll(logDir, os.ModeDir|os.ModePerm)
	}
	// logrus.SetReportCaller(true)
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	switch config.GetString("log.level") {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	}
	systemLogFilename := time.Now().In(time.FixedZone("CST", 8*3600)).Format("20060102-150405") + "-sys"
	systemLogPath := logDir + systemLogFilename + ".log"
	systemLogFile, err := os.OpenFile(systemLogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755) //只写、创建、追加
	if err != nil {
		fmt.Println("failed to open log, program exit")
		os.Exit(1)
	}
	mws := io.MultiWriter(os.Stdout, systemLogFile) // 同时输出到控制台和文件
	// mws := io.MultiWriter(systemLogFile) 			// 只输出到文件
	logrus.SetOutput(mws)
	logrus.Info("service started") // Panic/Fatal/Error/Warn/Info/Debug/Trace
	
	// # service
	app := fiber.New(fiber.Config{
		ServerHeader: "input-service",
		AppName:      "input service",
	})

	// # http log
	httpLogFilename := time.Now().In(time.FixedZone("CST", 8*3600)).Format("20060102-150405") + "-http"
	httpLogPath := logDir + httpLogFilename + ".log"
	httpLogFile, err := os.OpenFile(httpLogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755) //只写、创建、追加
	mwh := io.MultiWriter(os.Stdout, httpLogFile) 	// 同时输出到控制台和文件
	// mwh := io.MultiWriter(httpLogFile) 				// 只输出到文件
	app.Use(logger.New(logger.Config{
		Output: mwh,
	}))

	// # cors
	if config.GetBool("service.cors") == true {
		app.Use(cors.New())
		logrus.Info("cors on")
	}

	// # setup 
	// ## affair
	// err = affair.Setup(config.GetStringMap("affair"))
	// if err != nil {
	// 	logrus.Error("affair.Setup() ", err)
	// 	os.Exit(1)
	// }

	// # api
	api := app.Group("/api")

	api.Post("/testers/:testerId", func(c *fiber.Ctx) error {
		testerId := c.Params("testerId")
		var option map[string]string
		body := c.Body()
		err := json.Unmarshal(body, &option) 
		if err != nil { 
			logrus.Trace("post-/api/testers/:testerId-json.Unmarshal() ", err)
			return c.JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		err = trial.ChangeTo(testerId)
 		if err != nil {
			logrus.Trace("post-/api/testers/:testerId-trial.ChangeTo() ", err)
			return c.JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		go trial.Start(option)
		return c.JSON(&fiber.Map{
			"success": true,
		})
	}) 
  	api.Get("/datas/:testerId/:testId", func(c *fiber.Ctx) error {
		testerId := c.Params("testerId")
		testId := c.Params("testId")
		list, err := datas.ListRecord(testerId, testId)
		if err != nil {
			logrus.Trace("get-/api/datas/:testerId/:testId-datas.ListRecord() ", err)
			return c.JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		return c.JSON(list)
	})

	// # listen
	logrus.Fatal(app.Listen(config.GetString("service.host") + ":" + config.GetString("service.port")))
}
