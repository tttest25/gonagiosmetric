package nagiosclient

import (
	"net/http"
	"net/url"
	"strings"
)

func sendToNagios(str string) {
	form := url.Values{
		"perfdata": {"fed-serv;check_MISOJ_reception;0;OK - " + str},
	}
	form.Add("ln", "ln")
	form.Add("ip", "ip")
	form.Add("ua", "ua")

	req, err := http.NewRequest("POST", "http://10.59.20.16:8000", strings.NewReader(form.Encode()))
	if err != nil {
		Logger.Fatal(err)
	}
	req.SetBasicAuth("fed_monitor", "l3tm31n")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		Logger.Fatal(err)
	}
	Logger.Printf("Resp %#v  \n", resp)
}
