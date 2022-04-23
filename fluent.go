package domainpricescraper

import (
	"github.com/tunaitis/domainpricescraper/domain"
	"sort"
)

func (s *Scraper) UseDomain(domain string) *Scraper {
	s.Domains = append(s.Domains, domain)
	sort.Strings(s.Domains)
	return s
}

func (s *Scraper) UseAllDomains() *Scraper {
	s.Domains = domain.AllDomains()
	sort.Strings(s.Domains)
	return s
}

func (s *Scraper) UseRegistrar(r domain.Registrar) *Scraper {
	if r != nil {
		s.registrars = append(s.registrars, r)
	}
	return s
}
