package gandi

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tunaitis/domainpricescraper/util"
	"strconv"
	"strings"
)

type Gandi struct {
}

func New() *Gandi {
	return &Gandi{}
}

func (g Gandi) Name() string {
	return "Gandi"
}

func (g Gandi) Scrape(tld []string) (map[string]float64, error) {
	result := make(map[string]float64)

	docs := make(map[uint8]*goquery.Document)

	for i := range tld {
		l := tld[i][1]

		if docs[l] == nil {
			u := fmt.Sprintf("https://www.gandi.net/en-US/domain/tld?prefix=%c", l)
			doc, err := util.DownloadDocument(u, nil, nil)

			if err != nil {
				return nil, err
			}

			docs[l] = doc
		}

		doc := docs[l]

		doc.Find(".comparative-table tr").Each(func(j int, s *goquery.Selection) {
			anchor := s.Find("a")
			if anchor.Text() == tld[i] {
				prices := s.Find(".comparative-table__price")
				if prices.Length() == 0 {
					return
				}

				price := prices.First().Text()
				price = strings.Replace(price, "$", "", 1)

				var p float64
				p, err := strconv.ParseFloat(price, 64)
				if err != nil {
					return
				}

				result[tld[i]] = p

				return
			}
			return
		})
	}

	/*
		u := "https://www.gandi.net/en-US/domain/tld?prefix=c"
		_, err := util.DownloadString(u, nil)

		if err != nil {
			return nil, err
		}
	*/

	return result, nil
}
