package namecheap

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/tunaitis/domainpricescraper/util"
	"strconv"
	"strings"
)

type NameCheap struct {
	colly *colly.Collector
}

func New() *NameCheap {
	return &NameCheap{
		colly: colly.NewCollector(),
	}
}

func (n *NameCheap) Name() string {
	return "NameCheap"
}

func (n *NameCheap) getPrice(row *goquery.Selection, name string) (string, error) {
	col := row.Find(fmt.Sprintf("td[data-table-title=%s]", name))
	if col.Length() == 0 {
		return "", errors.New("column not found")
	}

	s := col.Children().First()
	if s.Length() == 0 {
		return "", errors.New("column contains no children")
	}

	return s.Text(), nil
}

func (n *NameCheap) Scrape(tld []string) (map[string]float64, error) {

	result := make(map[string]float64)

	n.colly.OnHTML(".gb-table", func(c *colly.HTMLElement) {
		c.ForEach(".gb-tld-name", func(i int, e *colly.HTMLElement) {

			if !util.Contains(tld, e.Text) {
				return
			}

			tr := e.DOM.ParentsFiltered("tr")
			if tr.Length() == 0 {
				return
			}

			price, err := n.getPrice(tr, "Register")
			if err != nil {
				return
			}

			price = strings.Replace(price, "$", "", 1)

			number, err := strconv.ParseFloat(price, 64)
			if err != nil {
				return
			}

			result[e.Text] = number
		})
	})

	err := n.colly.Visit("https://namecheap.com/domains/")
	if err != nil {
		return nil, err
	}

	return result, nil
}
