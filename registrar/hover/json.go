package hover

import "encoding/json"

type content struct {
	Prices map[string][]float64 `json:"prices"`
}

func parsePrice(data string) (map[string][]float64, error) {
	var c content
	err := json.Unmarshal([]byte(data), &c)
	if err != nil {
		return nil, err
	}
	return c.Prices, nil
}
