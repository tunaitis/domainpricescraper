package domainpricescraper

import (
	"github.com/gocolly/colly/v2"
	"github.com/tunaitis/domainpricescraper/domain"
)

type Scraper struct {
	registrars []domain.Registrar
	domains []string
	collector *colly.Collector
}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (s *Scraper) Scrape() (map[domain.Registrar]map[string]float64, error) {
	result := make(map[domain.Registrar]map[string]float64)

	for i := range s.registrars {
		e := s.registrars[i]

		r, err := e.Scrape(s.domains)
		if err != nil {
			return nil, err
		}

		result[e] = r
	}

	return result, nil
}