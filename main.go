package main

import (
	"github.com/tttest25/gonagiosmetric/logger"
	"github.com/tttest25/gonagiosmetric/nagiosclient"
	"github.com/tttest25/gonagiosmetric/scrapper"
	//"encoding/json"
)

var (
// Logger variable for logging
// l *log.Logger
// log nage

)

func main() {
	l := logger.ReturnLogger("main")
	l.Printf("Start ")
	strMetricOut := scrapper.Scrape()
	l.Printf("%s", strMetricOut)
	l.Printf("Elapsed %dms \n", logger.TimeElapsed()/1000)
	nagiosclient.SendToNagios(strMetricOut)

	l.Printf("Successfully stop")

}
