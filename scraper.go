package domainpricescraper

import (
	"github.com/gocolly/colly/v2"
	"github.com/tunaitis/domainpricescraper/domain"
)

type Scraper struct {
	registrars []domain.Registrar
	domains    []string
	collector  *colly.Collector
}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (s *Scraper) Scrape() (map[domain.Registrar]map[string]float64, error) {
	result := make(map[domain.Registrar]map[string]float64)

	for i := range s.registrars {
		e := s.registrars[i]

		// make a copy of the domain list before passing it to a registrar
		// to prevent the registrar from modifying the original slice
		d := make([]string, len(s.domains))
		copy(d, s.domains)

		r, err := e.Scrape(d)
		if err != nil {
			return nil, err
		}

		result[e] = r
	}

	return result, nil
}
