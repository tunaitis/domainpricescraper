package dynadot

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/tunaitis/domainpricescraper/util"
	"strconv"
	"strings"
)

type Dynadot struct {
	colly *colly.Collector
}

func New() *Dynadot {
	return &Dynadot{
		colly: colly.NewCollector(),
	}
}

func (d Dynadot) Name() string {
	return "Dynadot"
}

func (d Dynadot) Scrape(tld []string) (map[string]float64, error) {
	result := make(map[string]float64)

	doc, err := util.DownloadDocument("https://www.dynadot.com/change_currency.html?chg_currency=USD&pg=%2Fdomain%2Ftlds.html", nil, nil)
	if err != nil {
		return nil, err
	}

	table := doc.Find(".tld-table")
	if table.Length() == 0 {
		return nil, errors.New("table element not found")
	}

	table.Find(".tld-content").Each(func(i int, s *goquery.Selection) {
		anchor := s.Find("a")
		span := s.Find("span")

		if anchor.Length() == 0 || span.Length() == 0 {
			return
		}

		domain := anchor.Text()
		price := span.Text()
		price = strings.Replace(price, "$", "", 1)

		var p float64
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			return
		}

		if !util.Contains(tld, domain) {
			return
		}

		result[domain] = p
	})

	return result, nil
}
