package godaddy

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Price  int `json:"price"`
	Period int `json:"period"`
}

func parseApiResponse(data []byte) (float64, error) {
	var r response
	err := json.Unmarshal(data, &r)
	if err != nil {
		return 0, fmt.Errorf("%w (%s)", err, string(data))
	}

	return float64(r.Price/r.Period) / 1000000, nil
}
