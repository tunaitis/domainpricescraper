package googledomains

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/tunaitis/domainpricescraper/util"
	"strconv"
	"strings"
)

type GoogleDomains struct {
	colly *colly.Collector
}

func (g GoogleDomains) Name() string {
	return "GoogleDomains"
}

func (g GoogleDomains) Scrape(tld []string) (map[string]float64, error) {

	result := make(map[string]float64)

	g.colly.OnHTML(".nice-table.narrow", func(c *colly.HTMLElement) {
		c.ForEach("tr", func(i int, e *colly.HTMLElement) {

			var d string
			var p float64

			e.DOM.Find("td").Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					d = s.Text()
				}

				if i == 1 {
					t := s.Text()
					t = strings.Replace(t, "$", "", 1)
					t = strings.Replace(t, "USD", "", 1)
					t = strings.Trim(t, " ")

					n, err := strconv.ParseFloat(t, 64)
					if err != nil {
						return
					}

					p = n
				}
			})

			if !util.Contains(tld, d) {
				return
			}

			result[d] = p
		})
	})

	err := g.colly.Visit("https://support.google.com/domains/answer/6010092")
	if err != nil {
		return nil, err
	}

	return result, nil
}

func New() *GoogleDomains {
	return &GoogleDomains{
		colly: colly.NewCollector(),
	}
}
