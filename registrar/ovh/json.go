package ovh

import "encoding/json"

type root struct {
	OvhDomainNames ovhDomainNames `json:"ovh_domain_names"`
	//Prices map[string][]float64 `json:"prices"`
}

type ovhDomainNames struct {
	PricesData map[string]pricesData `json:"pricesData"`
}

type pricesData struct {
	Price float64 `json:"data-price"`
}

func parseJson(data string) (map[string]pricesData, error) {
	var r root
	err := json.Unmarshal([]byte(data), &r)
	if err != nil {
		return nil, err
	}
	return r.OvhDomainNames.PricesData, nil
}
