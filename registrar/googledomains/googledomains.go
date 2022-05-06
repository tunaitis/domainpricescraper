package googledomains

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/tunaitis/domainpricescraper/util"
	"strconv"
	"strings"
)

type GoogleDomains struct {
}

func New() *GoogleDomains {
	return &GoogleDomains{}
}

func (g GoogleDomains) Name() string {
	return "Google Domains"
}

func (g GoogleDomains) Scrape(tld []string) (map[string]float64, error) {

	result := make(map[string]float64)

	doc, err := util.DownloadDocument("https://support.google.com/domains/answer/6010092", nil, nil)
	if err != nil {
		return nil, err
	}

	table := doc.Find(".nice-table.narrow")
	if table.Length() == 0 {
		return nil, errors.New("table element not found")
	}

	table.Find("tr").Each(func(i int, s *goquery.Selection) {
		var d string
		var p float64

		s.Find("td").Each(func(i int, s *goquery.Selection) {
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

	return result, nil
}
