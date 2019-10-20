package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/goarne/config"
	"github.com/goarne/logging"
	"github.com/goarne/web"
)

var (
	cnf appConfig
)

type appConfig struct {
	ServerPort  int32 `yaml:"serviceurl"`
	Infologger  logging.LogConfig
	ErrorLogger logging.LogConfig
}

func init() {
	cnf = appConfig{}

	cl := config.ConfigLoader{"config.d/appconfig.yaml"}

	//Adding the concrete unmarshalling of AppConfig structure to the generic Config function.

	cl.LoadAppKonfig(&cnf)

	rotatingTraceWriter := logging.CreateRotatingWriter(cnf.Infologger)
	rotatingErrorWriter := logging.CreateRotatingWriter(cnf.ErrorLogger)

	tracerLogger := logging.CreateLogWriter(rotatingTraceWriter)
	tracerLogger.Append(os.Stdout)

	errorLogger := logging.CreateLogWriter(rotatingErrorWriter)
	errorLogger.Append(os.Stdout)

	logging.InitLoggers(tracerLogger, tracerLogger, errorLogger, errorLogger)
}

func main() {

	router := createWebRouter()

	if err := http.ListenAndServe(":"+strconv.FormatInt(8080, 10), router); err != nil {
		logging.Error.Println(err)
	}

	logging.Trace.Println("Server stopped.")
}

func createWebRouter() *web.WebRouter {
	rootPath := web.NewRoute().Path("/").Method(web.HttpGet).HandlerFunc(httpGetSample)

	router := web.NewWebRouter()
	router.AddRoute(rootPath)

	return router
}

func httpGetSample(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Sample JSON payload\n"))
}
