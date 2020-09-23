package nagiosclient

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tttest25/gonagiosmetric/logger"
)

var (
	// Logger variable for logging
	l *log.Logger
)

// SendToNagios - send http post to passive nagios server
func SendToNagios(str string) {
	form := url.Values{
		"perfdata": {"fed-serv;check_MISOJ_reception;0;OK - " + str},
	}
	form.Add("ln", "ln")
	form.Add("ip", "ip")
	form.Add("ua", "ua")

	req, err := http.NewRequest("POST", "http://10.59.20.16:8000", strings.NewReader(form.Encode()))
	if err != nil {
		l.Fatal(err)
	}
	req.SetBasicAuth("fed_monitor", "l3tm31n")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	cli := &http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := cli.Do(req)
	if err != nil {
		l.Fatal(err)
	}
	l.Printf("Resp %#v  \n", resp)
}

func init() {
	l = logger.ReturnLogger("nagios")

}
