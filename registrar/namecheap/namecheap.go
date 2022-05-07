package namecheap

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tunaitis/domainpricescraper/util"
	"strconv"
	"strings"
)

type NameCheap struct {
}

func New() *NameCheap {
	return &NameCheap{}
}

func (n *NameCheap) Name() string {
	return "namecheap"
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

	doc, err := util.DownloadDocument("https://namecheap.com/domains/", nil, nil)
	if err != nil {
		return nil, err
	}

	table := doc.Find(".gb-table")
	if table.Length() == 0 {
		return nil, errors.New("table element not found")
	}

	table.Find(".gb-tld-name").Each(func(i int, s *goquery.Selection) {
		t := s.Text()

		if !util.Contains(tld, t) {
			return
		}

		tr := s.ParentsFiltered("tr")
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

		result[t] = number
	})

	return result, nil
}
