package hover

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/tunaitis/domainpricescraper/util"
	"log"
	"strings"
)

type Hover struct {
	colly *colly.Collector
}

func New() *Hover {
	return &Hover{
		colly: colly.NewCollector(),
	}
}

func (h Hover) Name() string {
	return "Hover"
}

func (h Hover) Scrape(tld []string) (map[string]float64, error) {
	result := make(map[string]float64)

	for i := range tld {
		tld[i] = strings.Replace(tld[i], ".", "", 1)
	}

	h.colly.OnHTML("div.table", func(c *colly.HTMLElement) {
		prices, err := parsePrice(c.Attr("data-props"))
		if err != nil {
			log.Fatal(err)
			return
		}

		for k, v := range prices {
			if !util.Contains(tld, k) {
				continue
			}
			result[fmt.Sprintf(".%s", k)] = v[1]
		}
	})

	err := h.colly.Visit("https://www.hover.com/domain-pricing")
	if err != nil {
		return nil, err
	}

	return result, nil
}
