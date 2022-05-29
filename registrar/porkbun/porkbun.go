package porkbun

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tunaitis/domainpricescraper/util"
	"strconv"
	"strings"
)

type Porkbun struct {
}

func New() *Porkbun {
	p := &Porkbun{}
	return p
}

func (p Porkbun) Name() string {
	return "porkbun"
}

func (p Porkbun) Scrape(tld []string) (map[string]float64, error) {
	result := make(map[string]float64)

	doc, err := util.DownloadDocument("https://porkbun.com/products/domains", nil, nil)
	if err != nil {
		return nil, err
	}

	doc.Find(".domainsPricingAllExtensionsItem").Each(func(i int, s *goquery.Selection) {

		domainTag := s.Find("a")
		if domainTag.Length() == 0 {
			return
		}

		domain := domainTag.First().Text()
		if !util.Contains(tld, domain) {
			return
		}

		priceTag := s.Find("span.sortValue")
		if priceTag.Length() == 0 {
			return
		}

		priceText := priceTag.First().Text()
		priceText = strings.Replace(priceText, "$", "", 1)

		priceValue, err := strconv.ParseFloat(priceText, 64)
		if err != nil {
			return
		}

		result[domain] = priceValue
	})

	return result, nil
}
