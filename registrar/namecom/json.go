package namecom

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type price struct {
	Tld   string  `json:"tld"`
	Price float64 `json:"registration_price,string"`
}

func parsePrice(data []byte) ([]price, error) {
	// put quotes around every number in the json
	// because sometimes the json comes with quoted numbers and sometimes not
	re := regexp.MustCompile(`(":\s*)([\d\.]+)(\s*[,}])`)
	data = re.ReplaceAll(data, []byte(`$1"$2"$3`))

	var prices []price
	err := json.Unmarshal(data, &prices)
	if err != nil {
		return nil, fmt.Errorf("%w (%s)", err, string(data))
	}

	return prices, nil
}
