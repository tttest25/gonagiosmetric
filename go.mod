module github.com/tttest25/gonagiosmetric

 replace github.com/tttest25/gonagiosmetric/logger => ./logger
 replace github.com/tttest25/gonagiosmetric/nagiosclient => ./nagiosclient
 replace github.com/tttest25/gonagiosmetric/scrapper => ./scrapper

go 1.14

//require github.com/gocolly/colly/v2 v2.1.0

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/andybalholm/cascadia v1.2.0 // indirect
	github.com/tttest25/gonagiosmetric/logger v0.0.0-00010101000000-000000000000
	github.com/tttest25/gonagiosmetric/nagiosclient v0.0.0-00010101000000-000000000000
	github.com/tttest25/gonagiosmetric/scrapper v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
)
