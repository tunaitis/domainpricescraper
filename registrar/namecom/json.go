package namecom

import "encoding/json"

type price struct {
	Tld   string  `json:"tld"`
	Price float64 `json:"registration_price"`
}

func parsePrice(data []byte) ([]price, error) {
	var prices []price
	err := json.Unmarshal(data, &prices)
	if err != nil {
		return nil, err
	}

	return prices, nil
}
