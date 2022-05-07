package hover

import (
	"errors"
	"fmt"
	"github.com/tunaitis/domainpricescraper/util"
	"net/http"
	"strings"
)

type Hover struct {
	transport *http.Transport
}

func New(options ...Option) *Hover {
	h := &Hover{}

	for _, opt := range options {
		opt(h)
	}

	return h
}

func (h Hover) Name() string {
	return "hover"
}

func (h Hover) Scrape(tld []string) (map[string]float64, error) {
	result := make(map[string]float64)

	for i := range tld {
		tld[i] = strings.Replace(tld[i], ".", "", 1)
	}

	doc, err := util.DownloadDocument("https://www.hover.com/domain-pricing", nil, h.transport)
	if err != nil {
		return nil, err
	}

	table := doc.Find("div.table")
	if table.Length() == 0 {
		return nil, errors.New("table element not found")
	}

	priceAttr, _ := table.Attr("data-props")

	prices, err := parsePrice(priceAttr)
	if err != nil {
		return nil, err
	}

	for k, v := range prices {
		if !util.Contains(tld, k) {
			continue
		}
		result[fmt.Sprintf(".%s", k)] = v[1]
	}

	return result, nil
}
