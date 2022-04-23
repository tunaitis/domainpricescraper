package domainpricescraper_test

import (
	"github.com/tunaitis/domainpricescraper"
	"github.com/tunaitis/domainpricescraper/domain"
	"sort"
	"testing"
)

func TestScraper_UseDomain_DomainOrder(t *testing.T) {
	s := domainpricescraper.NewScraper()
	s.UseDomain(domain.Eu)
	s.UseDomain(domain.Dev)
	s.UseDomain(domain.Com)

	wanted := make([]string, len(s.Domains))
	copy(wanted, s.Domains)
	sort.Strings(wanted)

	for i := range wanted {
		if s.Domains[i] != wanted[i] {
			t.Errorf("got %s wanted %s", s.Domains, wanted)
			break
		}
	}
}

func TestScraper_UseAllDomains_DomainOrder(t *testing.T) {
	s := domainpricescraper.NewScraper()
	s.UseAllDomains()

	wanted := make([]string, len(s.Domains))
	copy(wanted, s.Domains)
	sort.Strings(wanted)

	for i := range wanted {
		if s.Domains[i] != wanted[i] {
			t.Errorf("got %s wanted %s", s.Domains, wanted)
			break
		}
	}
}
