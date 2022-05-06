package ovh

import (
	"errors"
	"github.com/tunaitis/domainpricescraper/util"
)

type OVH struct {
}

func New() *OVH {
	return &OVH{}
}

func (g OVH) Name() string {
	return "OVHcloud"
}

func (g OVH) Scrape(tld []string) (map[string]float64, error) {

	result := make(map[string]float64)

	doc, err := util.DownloadDocument("https://www.ovhcloud.com/en/domains/tld/", nil, nil)
	if err != nil {
		return nil, err
	}

	data := doc.Find("script[data-drupal-selector=\"drupal-settings-json\"]")
	if data.Length() == 0 {
		return nil, errors.New("script element containing the price data not found")
	}

	domains, err := parseJson(data.Text())
	if err != nil {
		return nil, err
	}

	for _, t := range tld {
		for k := range domains {
			if t[1:] == k {
				result[t] = domains[k].Price
			}
		}
	}

	return result, nil
}
