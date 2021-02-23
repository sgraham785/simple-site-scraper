# Simple Site Scraper

This is project scrapes everthing in `<html>` and follows links for the given site. Writing the html content to `./scrapes/<domain>/<scraped_page>.html`

## Install 
`go get -u github.com/sgraham785/simple-site-scrape`

## Usage
```
Usage of simple-site-scraper:
  -o string
        Path for scraped outputs (default "./scrapes")
  -s string
        Site domain to scrape
```

## Example
`simple-site-scraper -s example.com -o /path/to/output/scrapes`