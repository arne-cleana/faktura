package main

import (
	"fmt"
	"os"

	"github.com/goarne/config"
	"github.com/goarne/logging"
)

var (
	appConfig       AppConfig
	msortWebService Service
)

type AppConfig struct {
	ServerPort  int32 `yaml:"serviceurl"`
	Infologger  logging.LogConfig
	ErrorLogger logging.LogConfig
}

func init() {
	appConfig = AppConfig{}

	cl := config.ConfigLoader{"config.d/appconfig.yaml"}

	//Adding the concrete unmarshalling of AppConfig structure to the generic Config function.

	cl.LoadAppKonfig(&appConfig)

	rotatingTraceWriter := logging.CreateRotatingWriter(appConfig.Infologger)
	rotatingErrorWriter := logging.CreateRotatingWriter(appConfig.ErrorLogger)

	tracerLogger := logging.CreateLogWriter(rotatingTraceWriter)
	tracerLogger.Append(os.Stdout)

	errorLogger := logging.CreateLogWriter(rotatingErrorWriter)
	errorLogger.Append(os.Stdout)

	logging.InitLoggers(tracerLogger, tracerLogger, errorLogger, errorLogger)
}

func main() {
	HelloWorld()
}

//HelloWorld prints welcome message
func HelloWorld() {
	fmt.Printf("Hello world")
}
