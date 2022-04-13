package dynadot

import (
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

	d.colly.OnHTML(".tld-table", func(c *colly.HTMLElement) {
		c.ForEach(".tld-content", func(i int, e *colly.HTMLElement) {
			a := e.DOM.Find("a")
			s := e.DOM.Find("span")

			if a.Length() == 0 || s.Length() == 0 {
				return
			}

			domain := a.Text()
			price := s.Text()
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
	})

	err := d.colly.Visit("https://www.dynadot.com/change_currency.html?chg_currency=USD&pg=%2Fdomain%2Ftlds.html")
	if err != nil {
		return nil, err
	}

	return result, nil
}
