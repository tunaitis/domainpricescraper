package domainpricescraper

import "github.com/tunaitis/domainpricescraper/domain"

func (s *Scraper) UseDomain(domain string) *Scraper {
	s.domains = append(s.domains, domain)
	return s
}

func (s *Scraper) UseAllDomains() *Scraper {
	s.domains = domain.AllDomains()
	return s
}

func (s *Scraper) UseRegistrar(r domain.Registrar) *Scraper {
	if r != nil {
		s.registrars = append(s.registrars, r)
	}
	return s
}

