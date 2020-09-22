package main

import (
	"github.com/tttest25/nagios_metric_host_bash/logger"
	//"encoding/json"
)

var (
// Logger variable for logging
// l *log.Logger
// log nage

)

func main() {
	l := logger.ReturnLogger("main")
	l.Printf("Start")

	// strMetricOut := exampleScrape()
	l.Printf("Elapsed %dms \n", logger.TimeElapsed()/1000)
	// sendToNagios(strMetricOut)

	l.Printf("Stop")

}
