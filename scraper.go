package domainpricescraper

import (
	"github.com/tunaitis/domainpricescraper/domain"
)

type Scraper struct {
	registrars []domain.Registrar
	Domains    []string
}

type Result struct {
	Registrars map[string]map[string]float64
	Domains    map[string]map[string]float64
}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (s *Scraper) Scrape() (*Result, error) {

	//result := make(map[domain.Registrar]map[string]float64)

	result := &Result{
		Registrars: make(map[string]map[string]float64),
		Domains:    make(map[string]map[string]float64),
	}

	for i := range s.registrars {
		e := s.registrars[i]

		// make a copy of the domain list before passing it to a registrar
		// to prevent the registrar from modifying the original slice
		d := make([]string, len(s.Domains))
		copy(d, s.Domains)

		r, err := e.Scrape(d)
		if err != nil {
			return nil, err
		}

		result.Registrars[e.Name()] = r

		for k := range r {
			if result.Domains[k] == nil {
				result.Domains[k] = make(map[string]float64)
			}
			result.Domains[k][e.Name()] = r[k]
		}
	}

	return result, nil
}
