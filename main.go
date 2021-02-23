package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly"
)

const sitesDir = "sites/"

func main() {
	var siteURL string
	flag.StringVar(&siteURL, "s", "", "Site domain to scrape")
	flag.Parse()

	// Sanitize -s
	siteURL = strings.TrimPrefix(siteURL, "https://")
	siteURL = strings.TrimPrefix(siteURL, "http://")
	siteURL = strings.ReplaceAll(siteURL, "/", "")

	// Create Directories
	sitePath := filepath.Join(sitesDir, siteURL)
	os.MkdirAll(sitePath, os.ModePerm)

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(siteURL, "www."+siteURL),
	)

	c.OnHTML("html", func(e *colly.HTMLElement) {
		// Sanitize the file name
		var fName string = e.Request.URL.String()
		fName = strings.TrimPrefix(fName, "https://")
		fName = strings.TrimPrefix(fName, "http://")
		fName = strings.ReplaceAll(fName, "/", "_")

		// Get the HTML
		html, err := e.DOM.Html()
		err = ioutil.WriteFile(sitesDir+siteURL+"/"+fName+".html", []byte(html), 0644)
		if err != nil {
			panic(err)
		}
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on...
	c.Visit("https://" + siteURL + "/")
}
