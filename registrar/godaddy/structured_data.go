package godaddy

import "encoding/json"

type structuredData struct {
	Offers offers `json:"offers"`
}

type offers struct {
	Offers []offer `json:"offers"`
}

type offer struct {
	Name  string  `json:"alternateName"`
	Price float64 `json:"price,string"`
}

func parseStructuredData(input string) (*structuredData, error) {
	var d structuredData

	err := json.Unmarshal([]byte(input), &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
