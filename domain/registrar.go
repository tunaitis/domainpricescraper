package domain

type Registrar interface {
	Name() string
	Scrape(tld []string) (map[string]float64, error)
}
