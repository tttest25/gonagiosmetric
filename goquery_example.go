package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	//"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const logFilename string = "log.txt"

var (
	// Logger variable for logging
	Logger *log.Logger

	// log nage

)

// truncateFile copy file with skipping from head
func truncateFile(src string, lim int) {
	fin, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	fout, err := os.Create(src + "tmp")
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	// Offset is the number of bytes you want to exclude
	_, err = fin.Seek(int64(lim), io.SeekStart)
	if err != nil {
		panic(err)
	}

	n, err := io.Copy(fout, fin)
	fmt.Printf("Copied %d bytes, err: %v", n, err)

	if err := os.Remove(src); err != nil {
		panic(err)
	}

	if err := os.Rename(src+"tmp", src); err != nil {
		panic(err)
	}

}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func timeTrack(start time.Time) int64 {
	elapsed := time.Since(start)
	return elapsed.Nanoseconds() / 1000
}

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

func exampleScrape() string {
	// Request the HTML page.
	res, err := http.Get("https://reception.gorodperm.ru/index.php?id=280")
	if err != nil {
		Logger.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		Logger.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		Logger.Fatal(err)
	}

	strMetric := ""
	// Find the review items
	doc.Find("#stat").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//<div id="stat" data-database="0.0011 s" data-queries="7" data-application="0.0339 s" data-total="0.0349 s" data-source="database">stat</div>
		database, _ := s.Attr("data-database")
		queries, _ := s.Attr("data-queries")
		application, _ := s.Attr("data-application")
		total, _ := s.Attr("data-total")

		strMetric = fmt.Sprintf("Review %d: MODx DB %s - Queries %s - App %s - Total %s\n", i, database, queries, application, total)

	})
	return strMetric
}

func main() {

	fi, err := os.Stat(logFilename)
	if err != nil {
		log.Fatal(err)
	}
	// get the size
	size := fi.Size()

	if fileExists(logFilename) && size > 4000 {
		truncateFile(logFilename, 2000)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}

	start := time.Now()

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		Logger.Fatal(err)
	}
	Logger = log.New(file, "MAIN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.SetOutput(file)

	Logger.Printf("Start")

	strMetricOut := exampleScrape()
	Logger.Printf("Elapsed %dms \n", timeTrack(start)/1000)
	sendToNagios(strMetricOut)

	Logger.Printf("Stop")

}
