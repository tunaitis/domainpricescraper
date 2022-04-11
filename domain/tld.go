package domain

const (
	Com = ".com"
	Net = ".net"
	Org = ".org"
	Io  = ".io"
	Dev = ".dev"
	Eu  = ".eu"
	Xyz = ".xyz"
)

func AllDomains() []string {
	return []string{Com, Net, Org, Io, Dev, Eu, Xyz}
}
