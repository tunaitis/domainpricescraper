package godaddy

import (
	"fmt"
	"github.com/tunaitis/domainpricescraper/util"
	"io"
	"net/http"
	"strings"
)

type GoDaddy struct {
	transport *http.Transport
	apiKey    string
	apiSecret string
}

type ApiResponse struct {
	Price float64 ``
}

func New(options ...Option) *GoDaddy {
	g := &GoDaddy{}

	for _, opt := range options {
		opt(g)
	}

	return g
}

func (g *GoDaddy) Name() string {
	return "GoDaddy"
}

func (g *GoDaddy) getUsingApi(tld []string) (map[string]float64, error) {
	result := make(map[string]float64)

	h := map[string]string{
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("sso-key %s:%s", g.apiKey, g.apiSecret),
	}

	for _, t := range tld {
		d := fmt.Sprintf("401b30e3b8b5d629635a5c613cdb7919%s", t)
		u := fmt.Sprintf("https://api.godaddy.com/v1/domains/available?domain=%s&checkType=FAST&forTransfer=false", d)
		r, err := util.NewRequest(u, nil, h, nil)

		if err != nil {
			return nil, err
		}

		if r.StatusCode != http.StatusOK {
			resp, err := io.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("an error has occured: %s", string(resp))
		}

		resp, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		p, err := parseApiResponse(resp)
		if err != nil {
			return nil, err
		}

		result[t] = p
	}

	return result, nil
}

func (g *GoDaddy) getFromWebSite(tld []string) (map[string]float64, error) {
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
			return nil, fmt.Errorf("structured data not found, url: %s", u)
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

func (g *GoDaddy) Scrape(tld []string) (map[string]float64, error) {
	if g.apiKey != "" && g.apiSecret != "" {
		return g.getUsingApi(tld)
	}
	return g.getFromWebSite(tld)
}
