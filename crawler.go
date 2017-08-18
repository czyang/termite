package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var movieItems []MovieItem

func fetchOnePage(url string) string {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Referer", url)
	res, err := client.Do(req)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find(".grid-view .item .info ul").Each(func(i int, s *goquery.Selection) {
		// Find title
		title := s.Find("li a em").Text()
		// Find rate.
		rateStr := "0"
		s.Find("li").Each(func(j int, sj *goquery.Selection) {
			if j == 2 {
				rateStr, _ = sj.Find("span").Attr("class")
				switch rateStr {
				case "rating5-t":
					rateStr = "10"
				case "rating45-t":
					rateStr = "9"
				case "rating4-t":
					rateStr = "8"
				case "rating35-t":
					rateStr = "7"
				case "rating3-t":
					rateStr = "6"
				case "rating25-t":
					rateStr = "5"
				case "rating2-t":
					rateStr = "4"
				case "rating15-t":
					rateStr = "3"
				case "rating1-t":
					rateStr = "2"
				default:
					rateStr = "1"
				}
			}
		})
		// Rate time
		rateTime := s.Find("li .date").Text()
		// Find my review
		review := s.Find("li .comment").Text()
		// Find desc
		desc := s.Find(".intro").Text()
		// Find douban link
		doubanLink, _ := s.Find("li a").Attr("href")
		//fmt.Println(title, rateStr, review, desc, doubanLink)
		movieItems = append(movieItems, MovieItem{Title: title, MyRate: rateStr, RateTime: rateTime, MyReview: review,
			MovieDesc: desc, DoubanLink: doubanLink})
	})
	// Find next link
	nextPageLink, _ := doc.Find(".paginator .next link").Attr("href")
	return nextPageLink
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please specify Douban username.")
		os.Exit(1)
	}

	// Iterate pages
	baseURL := "https://movie.douban.com/people/" + args[0] + "/collect"
	nextPageUrl := fetchOnePage(baseURL)
	for nextPageUrl != "" {
		nextPageUrl = fetchOnePage(nextPageUrl)
	}

	// Create Excel file.
	createXLSX(movieItems)
}