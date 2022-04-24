package namecom

import (
	"encoding/json"
	"fmt"
)

type price struct {
	Tld   string  `json:"tld"`
	Price float64 `json:"registration_price"`
}

func parsePrice(data []byte) ([]price, error) {
	var prices []price
	err := json.Unmarshal(data, &prices)
	if err != nil {
		return nil, fmt.Errorf("%w (%s)", err, string(data))
	}

	return prices, nil
}
