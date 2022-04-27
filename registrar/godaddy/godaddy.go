package godaddy

import (
	"fmt"
	"github.com/tunaitis/domainpricescraper/util"
	"net/http"
	"strings"
)

type GoDaddy struct {
}

func New() *GoDaddy {
	return &GoDaddy{}
}

func (g GoDaddy) Name() string {
	return "GoDaddy"
}

func (g GoDaddy) Scrape(tld []string) (map[string]float64, error) {
	result := make(map[string]float64)

	for _, t := range tld {
		u := fmt.Sprintf("https://www.godaddy.com/tlds/%s-domain", t[1:])
		c := []*http.Cookie{
			&http.Cookie{
				Name:  "market",
				Value: "en-us",
			},
		}

		doc, err := util.DownloadDocument(u, c, nil)
		if err != nil {
			return nil, err
		}

		sd := doc.Find("script[type='application/ld+json']")
		if sd.Length() == 0 {
			continue
		}

		p, err := parseStructuredData(sd.Text())
		if err != nil {
			return nil, err
		}

		for _, x := range p.Offers.Offers {
			if strings.Contains(x.Name, "Registration") {
				result[t] = x.Price
			}
		}
	}

	return result, nil
}
