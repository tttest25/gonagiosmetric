package scrapper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/tttest25/gonagiosmetric/logger"
	//"encoding/json"
)

var (
	// Logger variable for logging
	l *log.Logger
)

// Scrape return data from http
func Scrape() string {
	l.Printf("Start scrapping")
	// Request the HTML page.
	res, err := http.Get("https://reception.gorodperm.ru/index.php?id=280")
	if err != nil {
		l.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		l.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		l.Fatal(err)
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
	l.Printf("get result")
	return strMetric
}

func init() {
	l = logger.ReturnLogger("scrapper")

}
