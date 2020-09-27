package main

import (
	"fmt"

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
	l.Printf("--- Start ")

	metric := scrapper.Scrape()

	l.Printf("Get metrics %s", metric.String())

	l.Printf("Elapsed scrape %dms \n", logger.TimeElapsed()/1000)

	// set service name

	nagiosclient.Ns.SetHost("fed-serv")
	nagiosclient.Ns.NagiosSetServiceName("check_MISOJ_reception")

	nagiosclient.Ns.NagiosAddParam(fmt.Sprintf("cur_chan: %d ", metric.Channel), 0)

	nagiosclient.Ns.NagiosAddPerformance(
		nagiosclient.NagiosPerfDataI("curChan", float64(metric.Channel), "", 5, 6, -1, -1),
	)

	nagiosclient.Ns.NagiosAddParam(fmt.Sprintf("src: %s ", metric.Source), 0)

	nagiosclient.Ns.NagiosAddService("db", metric.Database, "f", "s", 0.02, 0.5, -1, -1)
	nagiosclient.Ns.NagiosAddService("Queries", metric.Queries, "i", "", 9, 11, -1, -1)
	nagiosclient.Ns.NagiosAddService("App", metric.Application, "f", "s", 0.3, 0.5, -1, -1)
	nagiosclient.Ns.NagiosAddService("Total", metric.Total, "f", "s", 0.3, 0.5, -1, -1)

	nagiosclient.Ns.NagiosAddService("chan1", metric.Metrics[0], "f", "s", 0.5, 1, -1, -1)
	nagiosclient.Ns.NagiosAddService("chan2", metric.Metrics[1], "f", "s", 0.5, 1, -1, -1)

	//&scrapper.Metric{Database:0.013000000035390258, Queries:8, Application:0.21800000220537186,
	//	Total:0.23099999874830246, Metrics:[]float64{0.12487, 0.132427},
	//	Source:"database", Channel:2, Userip:"46.146.166.64", Proto:"https"}

	nagiosclient.Ns.AddUpdate()

	/*
		// add param1 database and logic of status
		status = nagiosclient.NagiosGetTresh(metric.Database, 0.02, 0.5)
		nagiosclient.Ns.NagiosAddParam(fmt.Sprintf("db: %f s", metric.Database), status)
		nagiosclient.Ns.NagiosAddPerformance(
			nagiosclient.NagiosPerfData("db", metric.Database, "s", 0.02, 0.5, -1, -1),
		)
	*/

	l.Printf("NagiosMetrics result -> \n%s%s", nagiosclient.Ns.NagiosPassive(), nagiosclient.Ns.NagiosOutput())

	// fed-serv;check_MISOJ_reception;0;OK - db: %f s, dbquery: %f , app: %f s, total: %f s
	nagiosclient.SendToNagios(
		fmt.Sprintf("%s%s", nagiosclient.Ns.NagiosPassive(), nagiosclient.Ns.NagiosOutput()),
	)

	l.Printf("Elapsed send to nagios %dms \n", logger.TimeElapsed()/1000)
	l.Printf("=== Successfully stop\n \n \n")

}
