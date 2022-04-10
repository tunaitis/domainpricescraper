package namecom

import (
	"fmt"
)

type NameCom struct {
}

func New() *NameCom {
	return &NameCom{}
}

func (n *NameCom) Name() string {
	return "Name"
}

func (n *NameCom) Scrape(tld []string) (map[string]float64, error) {
	result := make(map[string]float64)
	priceData, err := downloadPrices(tld)
	if err != nil {
		return nil, err
	}

	prices, err := parsePrice(priceData)
	if err != nil {
		return nil, err
	}

	for i := range prices {
		result[fmt.Sprintf(".%s", prices[i].Tld)] = prices[i].Price
	}

	return result, nil
}
